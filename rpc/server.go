// Copyright 2017 Square, Inc.

package rpc

import (
	"crypto/tls"
	"fmt"
	"net"
	"os"
	pb2 "mars_agent/pb"

	"github.com/go-cmd/cmd"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	health "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/status"

	marsCmd "mars_agent/cmd"
	"mars_agent/configs/base"
	. "mars_agent/configs/log"
	"mars_agent/service"
	"mars_agent/utils"
)

// A Server executes a whitelist of commands when called by clients.
type Server interface {
	// Start the gRPC server, non-blocking.
	StartServer() error

	// Stop the gRPC server gracefully.
	StopServer() error

	pb2.RCEAgentServer

	health.HealthServer
}

// Internal implementation of pb.RCEAgentServer interface.
type server struct {
	laddr      string           // host:port listen address
	tlsConfig  *tls.Config      // if secure
	whitelist  marsCmd.Runnable // commands from config file
	repo       marsCmd.Repo     // running commands
	grpcServer *grpc.Server     // gRPC server instance of this agent
}

// NewServer makes a new Server that listens on laddr and runs the whitelist
// of commands. If tlsConfig is nil, the sever is insecure.
func NewServer(laddr string, tlsConfig *tls.Config, whitelist marsCmd.Runnable) Server {
	// Set log flags here so other pkgs can't override in their init().
	//log.SetFlags(log.Ldate | log.Lmicroseconds | log.Lshortfile | log.LUTC)

	s := &server{
		laddr:     laddr,
		tlsConfig: tlsConfig,
		repo:      marsCmd.NewRepo(),
		whitelist: whitelist,
	}

	// Create a gRPC server and register this agent a implementing the
	// RCEAgentServer interface and protocol
	var grpcServer *grpc.Server
	if tlsConfig != nil {
		opt := grpc.Creds(credentials.NewTLS(tlsConfig))
		grpcServer = grpc.NewServer(opt)
	} else {
		grpcServer = grpc.NewServer()
	}
	s.grpcServer = grpcServer

	return s
}

func (s *server) StartServer() error {
	// Register the RCEAgent service with the gRPC server.
	pb2.RegisterRCEAgentServer(s.grpcServer, s)
	health.RegisterHealthServer(s.grpcServer, s)

	lis, err := net.Listen("tcp", s.laddr)
	if err != nil {
		return err
	}
	go s.grpcServer.Serve(lis)
	if s.tlsConfig != nil {
		//log.Printf("secure server listening on %s", s.laddr)
		Logger.Info("secure server listening on" + s.laddr)
	} else {
		Logger.Info("insecure server listening on" + s.laddr)
		//log.Printf("insecure server listening on %s", s.laddr)
	}
	return nil
}

func (s *server) StopServer() error {
	s.grpcServer.GracefulStop()
	Logger.Info("server stopped on" + s.laddr)

	//log.Printf("server stopped on %s", s.laddr)
	return nil
}

func (s *server) GorTraffic(ctx context.Context, command *pb2.Command) (*pb2.ID, error) {
	id := &pb2.ID{}
	Logger.Sugar().Info("receive command:", command)
	taskId := command.TaskId
	spec, err := s.whitelist.FindByName(command.Name)
	if err != nil {
		Logger.Warn("unknown command:" + command.Name)
		return id, status.Errorf(codes.InvalidArgument, "unknown command: %s", command.Name)
	}
	// 判断文件目录
	fileName := service.GorFileName(taskId)
	argus := []string{"--output-file", base.AppConfig.GetLogDir() + fileName, "--output-file-append"}
	// combined argus
	strings := append(spec.Args(), append(command.Arguments, argus...)...)
	userCmd := marsCmd.NewCmd(spec, strings)

	if err := s.repo.Add(userCmd); err != nil {
		Logger.Warn("duplicate command:" + fmt.Sprintf("%v", userCmd))
		return id, status.Errorf(codes.AlreadyExists, "duplicate command: %s", userCmd.Id)
	}
	userCmd.Cmd.Start()
	id.ID = userCmd.Id
	return id, nil
}

func (s *server) TrafficUpload(ctx context.Context, command *pb2.Command) (*pb2.ID, error) {
	id := &pb2.ID{}
	taskId := command.TaskId
	userCmd := &marsCmd.Cmd{
		Id:   marsCmd.Id(),
		Name: "task_oss_" + taskId,
		Args: []string{taskId},
	}
	if err := s.repo.Add(userCmd); err != nil {
		// This should never happen
		//log.Printf("duplicate command: %+v", marsCmd)
		Logger.Warn("duplicate command:" + fmt.Sprintf("%v", userCmd))
		return id, status.Errorf(codes.AlreadyExists, "duplicate command: %s", userCmd.Id)
	}
	//log
	Logger.Sugar().Info("receive traffic upload task, taskId:" + taskId)
	go func() {
		// 处理分析任务
		err := service.TrafficDyeingToCsv(base.AppConfig.GetLogDir(), taskId)
		if err != nil {
			userCmd.Complete = true
			userCmd.Error = err
			return
		}
		// 上传任务
		fileName := service.OssFileName(taskId)
		err = service.OssUpload(base.AppConfig.GetLogDir()+"/"+fileName, fileName)
		userCmd.Complete = true
		userCmd.Error = err
	}()
	id.ID = userCmd.Id
	return id, nil
}

func (s *server) Testing(ctx context.Context, command *pb2.Command) (*pb2.ID, error) {
	id := &pb2.ID{}
	Logger.Sugar().Info("receive command:", command)
	taskId := command.TaskId
	jmxId := command.JmxId
	spec, err := s.whitelist.FindByName(command.Name)
	if err != nil {
		Logger.Warn("unknown command:" + command.Name)
		return id, status.Errorf(codes.InvalidArgument, "unknown command: %s", command.Name)
	}
	userCmd := marsCmd.NewCmdWitchOutArgus(spec)
	if err := s.repo.Add(userCmd); err != nil {
		Logger.Warn("duplicate command:" + fmt.Sprintf("%v", userCmd))
		return id, status.Errorf(codes.AlreadyExists, "duplicate command: %s", userCmd.Id)
	}
	go func() {
		argus := []string{"-n"}
		argus = append(argus, command.Arguments...)
		jmxFilePath, err := checkJmxId(jmxId)
		if err != nil {
			userCmd.Error = err
			return
		}
		argus = append(argus, "-t", jmxFilePath)
		csvPath, err := checkCsvTaskId(taskId)
		if err != nil {
			userCmd.Error = err
			return
		}
		if csvPath != "" {
			argus = append(argus, "-JcsvData="+csvPath)
		}
		//argus := []string{"-n", "-t", "/Users/ruifeng/Downloads/jmeter/params.jmx", "-JcsvData=" + localFileName}
		userCmd.Args = append(userCmd.Args, argus...)
		userCmd.Cmd = cmd.NewCmd(spec.Path(), userCmd.Args...)
		userCmd.Cmd.Start()
	}()
	id.ID = userCmd.Id
	return id, nil
}

/**
上报用户的压测栅栏数据
*/
func (s *server) SampleTesting(ctx context.Context, command *pb2.Command) (*pb2.ID, error) {
	id := &pb2.ID{}
	Logger.Sugar().Info("receive sample testing command:", command)
	jmxId := command.JmxId
	spec, err := s.whitelist.FindByName(command.Name)
	if err != nil {
		Logger.Warn("unknown command:" + command.Name)
		return id, status.Errorf(codes.InvalidArgument, "unknown command: %s", command.Name)
	}
	userCmd := marsCmd.NewCmdWitchOutArgus(spec)
	if err := s.repo.Add(userCmd); err != nil {
		Logger.Warn("duplicate command:" + fmt.Sprintf("%v", userCmd))
		return id, status.Errorf(codes.AlreadyExists, "duplicate command: %s", userCmd.Id)
	}
	go func() {
		argus := []string{"-n"}
		var jmxFilePath string
		if command.SimpleTest {
			jmxFilePath, err = checkSampleJmxId(jmxId)
		} else {
			jmxFilePath, err = checkJmxId(jmxId)
		}

		if err != nil {
			userCmd.Error = err
			return
		}
		argus = append(argus, "-t", jmxFilePath)
		// 自定义压测脚本
		if !command.SimpleTest {
			argus = append(argus, command.Arguments...)
			userCmd.Args = append(userCmd.Args, argus...)
			userCmd.Cmd = cmd.NewCmd(spec.Path(), userCmd.Args...)
			Logger.Sugar().Info("jmeter run command: ", userCmd.Cmd.Args)
			userCmd.Cmd.Start()
			return
		}

		fileName := "jmeter_sample_output_" + jmxId + ".xml"
		path := base.AppConfig.GetLogDir()
		// 删除旧的栅栏文件
		err := os.Remove(path + fileName) //删除文件test.txt
		if err != nil {
			//如果删除失败则输出 file remove Error!
			Logger.Sugar().Error("删除文件失败: ", path+fileName, err)
		} else {
			//如果删除成功则输出 file remove OK!
			Logger.Sugar().Info(path + fileName + "删除文件成功: ")
		}
		// save jmeter sample output
		argus = append(argus, "-JsampleOutput="+path+fileName)
		//argus := []string{"-n", "-t", "/Users/ruifeng/Downloads/jmeter/params.jmx", "-JcsvData=" + localFileName}
		argus = append(argus, command.Arguments...)
		userCmd.Args = append(userCmd.Args, argus...)
		//options := cmd.Options{}

		userCmd.Cmd = cmd.NewCmd(spec.Path(), userCmd.Args...)
		//userCmd.Cmd = cmd.NewCmdOptions(options, spec.Path(), userCmd.Args...)
		Logger.Sugar().Info("sample run command: ", userCmd.Cmd.Args)
		statusChan := userCmd.Cmd.Start()
		// jmeter run block
		result := <-statusChan
		if result.Error != nil || result.Exit != 0 {
			Logger.Sugar().Error("sample test error, status:", result)
			return
		}
		// check file
		exist, err := utils.LocalFileExist(path + fileName)
		if err != nil {
			Logger.Sugar().Error("check local pathname error,", "filePath:"+path+fileName, err.Error())
			return
		}
		if !exist {
			Logger.Sugar().Error("sample output file not exist,", "filePath:"+path+fileName)
			return
		}
		// successful than upload sample output xml
		err = service.OssUpload(path+fileName, fileName)
		if err != nil {
			Logger.Sugar().Error("sample upload to oss error ,", "filePath:"+path+fileName, err.Error())
			return
		}
		Logger.Sugar().Info("sample output upload successful", "filePath:"+path+fileName)

	}()
	id.ID = userCmd.Id
	return id, nil
}

func checkCsvTaskId(taskId string) (string, error) {

	var taskOssName string
	if taskId == "" {
		// do not use traffic data
		return "", nil
	} else {
		// use outside traffic data
		taskOssName = service.OssFileName(taskId)
	}
	localCsvFileName := base.AppConfig.GetLogDir() + taskOssName
	// check local file exist
	fileExist, err := utils.LocalFileExist(localCsvFileName)
	if err != nil {
		err = fmt.Errorf("check local csv file error,filename:%s,err:%v", taskOssName, err)
		return "", err
	}
	if !fileExist {
		// check oss file is exist
		exits, err := service.OssObjectExist(taskOssName)
		if err != nil {
			err = fmt.Errorf("check oss csv file find error,filename:%s,err:%v", taskOssName, err)
			return "", err
		}
		if !exits {
			err = fmt.Errorf("csv file not exit in oss store,filename:%s", taskOssName)
			return "", err
		}
		err = service.OssDownloadFile(taskOssName, localCsvFileName)
		if err != nil {
			err = fmt.Errorf("download oss csv file error,name:%s,err:%v", taskOssName, err)
			return "", err
		}
	}
	return localCsvFileName, nil
}

func checkJmxId(jmxId string) (jmxFilePath string, err error) {
	var jmxName string
	if jmxId == "" {
		// default jmx
		jmxName = service.JmxFileName("default")
	} else {
		// use special jmx file
		jmxName = service.JmxFileName(jmxId)
	}
	jmxFilePath = base.AppConfig.GetLogDir() + jmxName
	// check local file exist
	//fileExist, err := utils.LocalFileExist(jmxFilePath)
	//if err != nil {
	//	err = fmt.Errorf("check local jmx file error,filename:%s,err:%v", jmxName, err)
	//	return "", err
	//
	//}
	//if !fileExist {
	//	// check oss file is exist
	//	exits, err := service.OssObjectExist(jmxName)
	//	if err != nil {
	//		err = fmt.Errorf("check oss jmx file find error,filename:%s,err:%v", jmxName, err)
	//		return "", err
	//	}
	//	if !exits {
	//		err = fmt.Errorf("jmx file not exit in oss store,filename:%s", jmxName)
	//		return "", err
	//	}
	//	err = service.OssDownloadFile(jmxName, jmxFilePath)
	//	if err != nil {
	//		err = fmt.Errorf("download oss jmx file error,name:%s,err:%v", jmxName, err)
	//		return "", err
	//	}
	//}
	//return jmxFilePath, nil
	exits, err := service.OssObjectExist(jmxName)
	if err != nil {
		err = fmt.Errorf("check oss jmx file find error,filename:%s,err:%v", jmxName, err)
		return "", err
	}
	if !exits {
		err = fmt.Errorf("jmx file not exit in oss store,filename:%s", jmxName)
		return "", err
	}
	err = service.OssDownloadFile(jmxName, jmxFilePath)
	if err != nil {
		err = fmt.Errorf("download oss jmx file error,name:%s,err:%v", jmxName, err)
		return "", err
	}
	return jmxFilePath, nil
}

func checkSampleJmxId(jmxId string) (jmxFilePath string, err error) {
	var jmxName string
	jmxName = service.JmxSampleFileName(jmxId)
	jmxFilePath = base.AppConfig.GetLogDir() + jmxName
	// check oss file exist
	exits, err := service.OssObjectExist(jmxName)
	if err != nil {
		err = fmt.Errorf("check oss jmx file find error,filename:%s,err:%v", jmxName, err)
		return "", err
	}
	if !exits {
		err = fmt.Errorf("jmx file not exit in oss store,filename:%s", jmxName)
		return "", err
	}
	err = service.OssDownloadFile(jmxName, jmxFilePath)
	if err != nil {
		err = fmt.Errorf("download oss jmx file error,name:%s,err:%v", jmxName, err)
		return "", err
	}
	return jmxFilePath, nil
}

// //////////////////////////////////////////////////////////////////////////
// pb.RCEAgentServer interface methods
// //////////////////////////////////////////////////////////////////////////

func (s *server) Start(ctx context.Context, c *pb2.Command) (*pb2.ID, error) {
	id := &pb2.ID{}

	spec, err := s.whitelist.FindByName(c.Name)
	if err != nil {
		//log.Printf("unknown command: %s", c.Name)
		Logger.Warn("unknown command:" + c.Name)
		return id, status.Errorf(codes.InvalidArgument, "unknown command: %s", c.Name)
	}

	// Append marsCmd request args to marsCmd spec args
	userCmd := marsCmd.NewCmd(spec, append(spec.Args(), c.Arguments...))
	if err := s.repo.Add(userCmd); err != nil {
		// This should never happen
		//log.Printf("duplicate command: %+v", marsCmd)
		Logger.Warn("duplicate command:" + fmt.Sprintf("%v", userCmd))
		return id, status.Errorf(codes.AlreadyExists, "duplicate command: %s", userCmd.Id)
	}
	Logger.Info(fmt.Sprintf("marsCmd=%s: start: %s path: %s args: %v", userCmd.Id, c.Name, spec.Path(), userCmd.Args))
	//log.Printf("marsCmd=%s: start: %s path: %s args: %v", marsCmd.Id, c.Name, spec.Path(), marsCmd.Args)
	userCmd.Cmd.Start()
	id.ID = userCmd.Id
	return id, nil
}

func (s *server) Check(context.Context, *health.HealthCheckRequest) (*health.HealthCheckResponse, error) {
	return &health.HealthCheckResponse{Status: health.HealthCheckResponse_SERVING}, nil
}

func (s *server) Watch(*health.HealthCheckRequest, health.Health_WatchServer) error {
	return nil

}

func (s *server) Wait(ctx context.Context, id *pb2.ID) (*pb2.Status, error) {
	//log.Printf("marsCmd=%s: wait", id.ID)
	Logger.Info(fmt.Sprintf("marsCmd=%s: wait", id.ID))
	//defer log.Printf("marsCmd=%s: wait return", id.ID)
	defer Logger.Info(fmt.Sprintf("marsCmd=%s: wait return", id.ID))

	userCms := s.repo.Get(id.ID)
	if userCms == nil {
		return nil, notFound(id)
	}
	// Reap the command
	defer s.repo.Remove(id.ID)

	// Wait for command or ctx to finish
	select {
	case <-userCms.Cmd.Done():
	case <-ctx.Done():
	}

	// Get final status of command and convert to pb.Status. If ctx was canceled
	// and command still running, its status will indicate this and ctx.Err()
	// below will return an error, else it will return nil.
	return mapStatus(userCms), ctx.Err()
}

func (s *server) GetStatus(ctx context.Context, id *pb2.ID) (*pb2.Status, error) {
	//log.Printf("marsCmd=%s: status", id.ID)
	Logger.Info(fmt.Sprintf("marsCmd=%s: status", id.ID))
	cmd := s.repo.Get(id.ID)
	if cmd == nil {
		return nil, notFound(id)
	}
	return mapStatus(cmd), nil
}

func (s *server) Stop(ctx context.Context, id *pb2.ID) (*pb2.Empty, error) {
	//log.Printf("marsCmd=%s: stop", id.ID)
	Logger.Info(fmt.Sprintf("marsCmd=%s: stop", id.ID))

	cmd := s.repo.Get(id.ID)
	if cmd == nil {
		return nil, notFound(id)
	}

	cmd.Cmd.Stop()

	return &pb2.Empty{}, nil
}

func (s *server) Running(empty *pb2.Empty, stream pb2.RCEAgent_RunningServer) error {
	//log.Println("list running")
	Logger.Info(fmt.Sprintf("list running"))
	for _, id := range s.repo.All() {
		if err := stream.Send(&pb2.ID{ID: id}); err != nil {
			return err
		}
	}
	return nil
}

func notFound(id *pb2.ID) error {
	return status.Errorf(codes.NotFound, "command ID %s not found", id.ID)
}

func mapStatus(cmd *marsCmd.Cmd) *pb2.Status {
	if cmd.Cmd == nil {
		return emptyCmd(cmd)
	}
	cmdStatus := cmd.Cmd.Status()

	var errMsg string
	if cmdStatus.Error != nil {
		errMsg = cmdStatus.Error.Error()
	}
	if cmd.Error != nil {
		errMsg += "\n" + cmd.Error.Error()
	}

	// Make a pb.Status struct by adding and mapping some fields
	pbStatus := &pb2.Status{
		ID:        cmd.Id,                // add
		Name:      cmd.Name,              // add
		ExitCode:  int64(cmdStatus.Exit), // map
		Error:     errMsg,                // map
		PID:       int64(cmdStatus.PID),  // map
		StartTime: cmdStatus.StartTs,     // map
		StopTime:  cmdStatus.StopTs,      // map
		Args:      cmd.Args,              // map
		Stdout:    cmdStatus.Stdout,      // same
		Stderr:    cmdStatus.Stderr,      // same
	}

	// Map go-marsCmd status to pb state
	switch {
	case cmdStatus.StartTs == 0 && cmdStatus.StopTs == 0:
		pbStatus.State = pb2.STATE_PENDING
	case cmdStatus.StartTs > 0 && cmdStatus.StopTs == 0:
		pbStatus.State = pb2.STATE_RUNNING
	case cmdStatus.StopTs > 0 && cmdStatus.Exit == 0:
		pbStatus.State = pb2.STATE_COMPLETE
	case cmdStatus.StopTs > 0 && cmdStatus.Exit != 0:
		pbStatus.State = pb2.STATE_FAIL
	default:
		pbStatus.State = pb2.STATE_UNKNOWN
	}

	return pbStatus
}

func emptyCmd(cmd *marsCmd.Cmd) *pb2.Status {
	var code int64
	errorMsg := ""
	var state pb2.STATE
	var stack []string
	if cmd.Complete || cmd.Error != nil {
		// 已经完成
		if cmd.Error != nil {
			errorMsg = cmd.Error.Error()
			state = pb2.STATE_FAIL
			stack = append(stack, utils.GetErrorStack(cmd.Error))
		} else {
			code = 0
			state = pb2.STATE_COMPLETE
		}
	} else {
		// 还未完成
		state = pb2.STATE_RUNNING
		code = -1
	}

	pbStatus := &pb2.Status{
		ID:        cmd.Id,   // add
		Name:      cmd.Name, // add
		ExitCode:  code,     // map
		State:     state,
		Error:     errorMsg,   // map
		PID:       int64(0),   // map
		StartTime: int64(0),   // map
		StopTime:  int64(0),   // map
		Args:      cmd.Args,   // map
		Stdout:    []string{}, // same
		Stderr:    stack,      // same
	}
	return pbStatus
}
