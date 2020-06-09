package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	. "mars_agent/configs/base"
	"mars_agent/configs/code"
	. "mars_agent/configs/log"
	"mars_agent/consul"
	"mars_agent/rpc"
	"strconv"
	"syscall"
)

var (
	h       bool
	s       string
	version bool
	Version string
	Build   string
	Tag     string
	nameMap = map[string]string{"rec": "mars_gor", "play": "mars_jmeter"}
)

func init() {
	flag.BoolVar(&version, "version", false, "print agent version information")
	flag.BoolVar(&h, "h", false, "this help")
	// 注意 `signal`。默认是 -s string，有了 `signal` 之后，变为 -s signal
	flag.StringVar(&s, "s", "", "start consul mode by `name` : gor, jmeter")
	//flag.StringVar(&p, "p", "/usr/local/nginx/", "set `prefix` path")
	// change default flag usage to print user defined cli help
	flag.Usage = usage
}

func usage() {
	_, _ = fmt.Fprintf(os.Stderr, ` mars agent version: %s
Usage: mars_agent  [-s name] [-version] [-h]

Options:
`, Version)
	flag.PrintDefaults()
}

func versionMessage() {
	fmt.Println(" mars agent \n" +
		"git_tag:" + Tag + ", version:" + Version + ", build_time: (" + Build + ")")
}

func main() {
	flag.Parse()
	if version {
		versionMessage()
		os.Exit(0)
	}
	name, ok := nameMap[s]
	if h {
		flag.Usage()
		os.Exit(0)
	}
	if s == "" || !ok {
		fmt.Printf("error -s paramas:%s , just only: rec, play \n", s)
		flag.Usage()
		os.Exit(code.ErrorFlag)
	}
	server := initAgent(name)

	c := make(chan os.Signal)
	// listener kill signal
	signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGUSR1, syscall.SIGUSR2)
	// block until receive kill signal
	s := <-c
	Logger.Sugar().Info(name + " agent service receive kill signal: " + s.String())
	// kill agent graceful
	err := server.StopServer()
	if err != nil {
		Logger.Sugar().Warn("logout "+name+" agent service failure!!!", err)
		return
	}
	Logger.Sugar().Info("logout " + name + " agent service successful!!!")
}

func initAgent(name string) rpc.Server {
	// init consul component
	consul.InitConsulConfig(name)
	// init log component
	InitLogConfig(AppConfig.GetLogDir(), name)
	// init grpc server
	server := rpc.NewServer("0.0.0.0:"+strconv.Itoa(AppConfig.App.Port), nil, AppConfig.Commands)
	err := server.StartServer()
	if err != nil {
		Logger.Sugar().Error("start "+name+" agent server failure!!!", err)
		os.Exit(code.GrpcStartError)
	}
	// grpc server register to consul
	err = consul.DoRegisterGrpcService(name, consul.LocalIp[0], AppConfig.App.Port)
	if err != nil {
		Logger.Sugar().Error("register "+name+" service to  consul failure!!!", err)
		os.Exit(code.GrpcRegisterError)
	}
	return server
}
