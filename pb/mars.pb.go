// Code generated by protoc-gen-go. DO NOT EDIT.
// source: mars.proto

package pb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type STATE int32

const (
	STATE_UNKNOWN  STATE = 0
	STATE_PENDING  STATE = 1
	STATE_RUNNING  STATE = 2
	STATE_COMPLETE STATE = 3
	STATE_FAIL     STATE = 4
	STATE_TIMEOUT  STATE = 5
	STATE_STOPPED  STATE = 6
)

var STATE_name = map[int32]string{
	0: "UNKNOWN",
	1: "PENDING",
	2: "RUNNING",
	3: "COMPLETE",
	4: "FAIL",
	5: "TIMEOUT",
	6: "STOPPED",
}

var STATE_value = map[string]int32{
	"UNKNOWN":  0,
	"PENDING":  1,
	"RUNNING":  2,
	"COMPLETE": 3,
	"FAIL":     4,
	"TIMEOUT":  5,
	"STOPPED":  6,
}

func (x STATE) String() string {
	return proto.EnumName(STATE_name, int32(x))
}

func (STATE) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_edc1aaf022af6c06, []int{0}
}

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_edc1aaf022af6c06, []int{0}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

type Status struct {
	ID                   string   `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	State                STATE    `protobuf:"varint,3,opt,name=State,proto3,enum=mars.STATE" json:"State,omitempty"`
	PID                  int64    `protobuf:"varint,4,opt,name=PID,proto3" json:"PID,omitempty"`
	StartTime            int64    `protobuf:"varint,5,opt,name=StartTime,proto3" json:"StartTime,omitempty"`
	StopTime             int64    `protobuf:"varint,6,opt,name=StopTime,proto3" json:"StopTime,omitempty"`
	ExitCode             int64    `protobuf:"varint,7,opt,name=ExitCode,proto3" json:"ExitCode,omitempty"`
	Args                 []string `protobuf:"bytes,8,rep,name=Args,proto3" json:"Args,omitempty"`
	Stdout               []string `protobuf:"bytes,9,rep,name=Stdout,proto3" json:"Stdout,omitempty"`
	Stderr               []string `protobuf:"bytes,10,rep,name=Stderr,proto3" json:"Stderr,omitempty"`
	Error                string   `protobuf:"bytes,11,opt,name=Error,proto3" json:"Error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Status) Reset()         { *m = Status{} }
func (m *Status) String() string { return proto.CompactTextString(m) }
func (*Status) ProtoMessage()    {}
func (*Status) Descriptor() ([]byte, []int) {
	return fileDescriptor_edc1aaf022af6c06, []int{1}
}

func (m *Status) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Status.Unmarshal(m, b)
}
func (m *Status) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Status.Marshal(b, m, deterministic)
}
func (m *Status) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Status.Merge(m, src)
}
func (m *Status) XXX_Size() int {
	return xxx_messageInfo_Status.Size(m)
}
func (m *Status) XXX_DiscardUnknown() {
	xxx_messageInfo_Status.DiscardUnknown(m)
}

var xxx_messageInfo_Status proto.InternalMessageInfo

func (m *Status) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *Status) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Status) GetState() STATE {
	if m != nil {
		return m.State
	}
	return STATE_UNKNOWN
}

func (m *Status) GetPID() int64 {
	if m != nil {
		return m.PID
	}
	return 0
}

func (m *Status) GetStartTime() int64 {
	if m != nil {
		return m.StartTime
	}
	return 0
}

func (m *Status) GetStopTime() int64 {
	if m != nil {
		return m.StopTime
	}
	return 0
}

func (m *Status) GetExitCode() int64 {
	if m != nil {
		return m.ExitCode
	}
	return 0
}

func (m *Status) GetArgs() []string {
	if m != nil {
		return m.Args
	}
	return nil
}

func (m *Status) GetStdout() []string {
	if m != nil {
		return m.Stdout
	}
	return nil
}

func (m *Status) GetStderr() []string {
	if m != nil {
		return m.Stderr
	}
	return nil
}

func (m *Status) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type ID struct {
	ID                   string   `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ID) Reset()         { *m = ID{} }
func (m *ID) String() string { return proto.CompactTextString(m) }
func (*ID) ProtoMessage()    {}
func (*ID) Descriptor() ([]byte, []int) {
	return fileDescriptor_edc1aaf022af6c06, []int{2}
}

func (m *ID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ID.Unmarshal(m, b)
}
func (m *ID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ID.Marshal(b, m, deterministic)
}
func (m *ID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ID.Merge(m, src)
}
func (m *ID) XXX_Size() int {
	return xxx_messageInfo_ID.Size(m)
}
func (m *ID) XXX_DiscardUnknown() {
	xxx_messageInfo_ID.DiscardUnknown(m)
}

var xxx_messageInfo_ID proto.InternalMessageInfo

func (m *ID) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

type Task struct {
	TaskId               int64    `protobuf:"varint,1,opt,name=TaskId,proto3" json:"TaskId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Task) Reset()         { *m = Task{} }
func (m *Task) String() string { return proto.CompactTextString(m) }
func (*Task) ProtoMessage()    {}
func (*Task) Descriptor() ([]byte, []int) {
	return fileDescriptor_edc1aaf022af6c06, []int{3}
}

func (m *Task) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Task.Unmarshal(m, b)
}
func (m *Task) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Task.Marshal(b, m, deterministic)
}
func (m *Task) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Task.Merge(m, src)
}
func (m *Task) XXX_Size() int {
	return xxx_messageInfo_Task.Size(m)
}
func (m *Task) XXX_DiscardUnknown() {
	xxx_messageInfo_Task.DiscardUnknown(m)
}

var xxx_messageInfo_Task proto.InternalMessageInfo

func (m *Task) GetTaskId() int64 {
	if m != nil {
		return m.TaskId
	}
	return 0
}

type Command struct {
	Name                 string   `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	TaskId               string   `protobuf:"bytes,2,opt,name=TaskId,proto3" json:"TaskId,omitempty"`
	JmxId                string   `protobuf:"bytes,3,opt,name=JmxId,proto3" json:"JmxId,omitempty"`
	SimpleTest           bool     `protobuf:"varint,4,opt,name=SimpleTest,proto3" json:"SimpleTest,omitempty"`
	Arguments            []string `protobuf:"bytes,5,rep,name=Arguments,proto3" json:"Arguments,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Command) Reset()         { *m = Command{} }
func (m *Command) String() string { return proto.CompactTextString(m) }
func (*Command) ProtoMessage()    {}
func (*Command) Descriptor() ([]byte, []int) {
	return fileDescriptor_edc1aaf022af6c06, []int{4}
}

func (m *Command) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Command.Unmarshal(m, b)
}
func (m *Command) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Command.Marshal(b, m, deterministic)
}
func (m *Command) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Command.Merge(m, src)
}
func (m *Command) XXX_Size() int {
	return xxx_messageInfo_Command.Size(m)
}
func (m *Command) XXX_DiscardUnknown() {
	xxx_messageInfo_Command.DiscardUnknown(m)
}

var xxx_messageInfo_Command proto.InternalMessageInfo

func (m *Command) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Command) GetTaskId() string {
	if m != nil {
		return m.TaskId
	}
	return ""
}

func (m *Command) GetJmxId() string {
	if m != nil {
		return m.JmxId
	}
	return ""
}

func (m *Command) GetSimpleTest() bool {
	if m != nil {
		return m.SimpleTest
	}
	return false
}

func (m *Command) GetArguments() []string {
	if m != nil {
		return m.Arguments
	}
	return nil
}

func init() {
	proto.RegisterEnum("mars.STATE", STATE_name, STATE_value)
	proto.RegisterType((*Empty)(nil), "mars.Empty")
	proto.RegisterType((*Status)(nil), "mars.Status")
	proto.RegisterType((*ID)(nil), "mars.ID")
	proto.RegisterType((*Task)(nil), "mars.Task")
	proto.RegisterType((*Command)(nil), "mars.Command")
}

func init() { proto.RegisterFile("mars.proto", fileDescriptor_edc1aaf022af6c06) }

var fileDescriptor_edc1aaf022af6c06 = []byte{
	// 541 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x53, 0xdd, 0x6a, 0xdb, 0x4c,
	0x10, 0xf5, 0xea, 0xc7, 0x92, 0xc7, 0x49, 0x10, 0x4b, 0x08, 0x8b, 0xf9, 0xc8, 0xe7, 0x88, 0x94,
	0x9a, 0x5c, 0x98, 0x92, 0x3e, 0x81, 0x6b, 0xa9, 0x46, 0x6d, 0x22, 0x0b, 0x49, 0x26, 0xd0, 0xab,
	0xca, 0xd1, 0xda, 0x88, 0x44, 0x3f, 0x5d, 0xad, 0x21, 0x7d, 0x83, 0xf6, 0xb2, 0x0f, 0xd9, 0xf7,
	0x28, 0xbb, 0x52, 0x6c, 0x51, 0x68, 0x72, 0xe5, 0x39, 0x67, 0x8e, 0x87, 0x99, 0xb3, 0x47, 0x00,
	0x79, 0xc2, 0xea, 0x69, 0xc5, 0x4a, 0x5e, 0x62, 0x4d, 0xd4, 0xb6, 0x01, 0xba, 0x9b, 0x57, 0xfc,
	0xbb, 0xfd, 0x4b, 0x81, 0x7e, 0xc4, 0x13, 0xbe, 0xab, 0xf1, 0x09, 0x28, 0x9e, 0x43, 0xd0, 0x18,
	0x4d, 0x06, 0xa1, 0xe2, 0x39, 0x18, 0x83, 0xe6, 0x27, 0x39, 0x25, 0x8a, 0x64, 0x64, 0x8d, 0x2f,
	0x40, 0x17, 0x6a, 0x4a, 0xd4, 0x31, 0x9a, 0x9c, 0x5c, 0x0f, 0xa7, 0x72, 0x72, 0x14, 0xcf, 0x62,
	0x37, 0x6c, 0x3a, 0xd8, 0x02, 0x35, 0xf0, 0x1c, 0xa2, 0x8d, 0xd1, 0x44, 0x0d, 0x45, 0x89, 0xff,
	0x83, 0x41, 0xc4, 0x13, 0xc6, 0xe3, 0x2c, 0xa7, 0x44, 0x97, 0xfc, 0x81, 0xc0, 0x23, 0x30, 0x23,
	0x5e, 0x56, 0xb2, 0xd9, 0x97, 0xcd, 0x3d, 0x16, 0x3d, 0xf7, 0x29, 0xe3, 0xf3, 0x32, 0xa5, 0xc4,
	0x68, 0x7a, 0xcf, 0x58, 0xac, 0x37, 0x63, 0xdb, 0x9a, 0x98, 0x63, 0x55, 0xac, 0x27, 0x6a, 0x7c,
	0x26, 0x8e, 0x49, 0xcb, 0x1d, 0x27, 0x03, 0xc9, 0xb6, 0xa8, 0xe5, 0x29, 0x63, 0x04, 0xf6, 0x3c,
	0x65, 0x0c, 0x9f, 0x82, 0xee, 0x32, 0x56, 0x32, 0x32, 0x94, 0x37, 0x36, 0xc0, 0x3e, 0x15, 0x46,
	0xfc, 0x6d, 0x87, 0x7d, 0x0e, 0x5a, 0x9c, 0xd4, 0x0f, 0x62, 0x96, 0xf8, 0xf5, 0x52, 0xd9, 0x53,
	0xc3, 0x16, 0xd9, 0x3f, 0x11, 0x18, 0xf3, 0x32, 0xcf, 0x93, 0x22, 0xdd, 0x5b, 0x87, 0x3a, 0xd6,
	0x1d, 0xfe, 0xd7, 0x18, 0xda, 0x22, 0xb1, 0xc3, 0xa7, 0xfc, 0xc9, 0x4b, 0xa5, 0xa5, 0x83, 0xb0,
	0x01, 0xf8, 0x1c, 0x20, 0xca, 0xf2, 0xea, 0x91, 0xc6, 0xb4, 0xe6, 0xd2, 0x4c, 0x33, 0xec, 0x30,
	0xc2, 0xd3, 0x19, 0xdb, 0xee, 0x72, 0x5a, 0xf0, 0x9a, 0xe8, 0xf2, 0xa8, 0x03, 0x71, 0xf5, 0x15,
	0x74, 0xf9, 0x26, 0x78, 0x08, 0xc6, 0xca, 0xff, 0xec, 0x2f, 0xef, 0x7c, 0xab, 0x27, 0x40, 0xe0,
	0xfa, 0x8e, 0xe7, 0x2f, 0x2c, 0x24, 0x40, 0xb8, 0xf2, 0x7d, 0x01, 0x14, 0x7c, 0x04, 0xe6, 0x7c,
	0x79, 0x1b, 0xdc, 0xb8, 0xb1, 0x6b, 0xa9, 0xd8, 0x04, 0xed, 0xe3, 0xcc, 0xbb, 0xb1, 0x34, 0x21,
	0x8a, 0xbd, 0x5b, 0x77, 0xb9, 0x8a, 0x2d, 0x5d, 0x80, 0x28, 0x5e, 0x06, 0x81, 0xeb, 0x58, 0xfd,
	0xeb, 0xdf, 0x0a, 0x98, 0xe1, 0xdc, 0x9d, 0x6d, 0x69, 0xc1, 0xb1, 0x2d, 0x53, 0xc1, 0x38, 0x3e,
	0x6e, 0xf2, 0xd0, 0xda, 0x30, 0x32, 0x1b, 0xe8, 0x39, 0x76, 0x0f, 0x8f, 0x41, 0xbb, 0x4b, 0x32,
	0x8e, 0xf7, 0xdc, 0xe8, 0xa8, 0x0d, 0x8f, 0x4c, 0x9f, 0xdd, 0xc3, 0x6f, 0x60, 0xb0, 0xa0, 0xbc,
	0x0d, 0xe3, 0xbf, 0x65, 0xff, 0x83, 0x26, 0xf2, 0xd1, 0x51, 0xb4, 0x29, 0x6c, 0x02, 0xdd, 0xc3,
	0x6f, 0x01, 0x16, 0x25, 0x8b, 0x59, 0xb2, 0xd9, 0x64, 0xf7, 0x2f, 0xad, 0x74, 0x05, 0xc7, 0xad,
	0x6a, 0x55, 0x3d, 0x96, 0x49, 0xfa, 0x92, 0xf6, 0x12, 0x0c, 0xe1, 0x7b, 0x56, 0x6c, 0x5f, 0x99,
	0x18, 0x25, 0xcf, 0x6f, 0xf4, 0x8a, 0xf6, 0x12, 0x8c, 0x70, 0x57, 0x14, 0x42, 0xd5, 0x3d, 0xa0,
	0xab, 0x79, 0x87, 0x3e, 0x5c, 0xc0, 0xd9, 0x7d, 0x99, 0x4f, 0xeb, 0x6f, 0x0f, 0xeb, 0xe9, 0xe1,
	0x2b, 0x5e, 0xef, 0x36, 0x01, 0xfa, 0xa2, 0x54, 0xeb, 0x1f, 0x08, 0xad, 0xfb, 0x92, 0x7a, 0xff,
	0x27, 0x00, 0x00, 0xff, 0xff, 0xb8, 0x8c, 0x3a, 0xc9, 0xe6, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RCEAgentClient is the client API for RCEAgent service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RCEAgentClient interface {
	// Start a command and immediately return its ID. Be sure to call Wait or Stop
	// to reap the command, else the agent will effectively leak memory by holding
	// unreaped commands. A command is considered running until reaped.
	Start(ctx context.Context, in *Command, opts ...grpc.CallOption) (*ID, error)
	// Wait for a command to complete or be stopped, reap it, and return its final status.
	Wait(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Status, error)
	// Get the status of a command if it hasn't been reaped by calling Wait or Stop.
	GetStatus(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Status, error)
	// Stop then reap a command by sending it a SIGTERM signal.
	Stop(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Empty, error)
	// start gor task
	GorTraffic(ctx context.Context, in *Command, opts ...grpc.CallOption) (*ID, error)
	// upload traffic to oss
	TrafficUpload(ctx context.Context, in *Command, opts ...grpc.CallOption) (*ID, error)
	// start jmeter task
	Testing(ctx context.Context, in *Command, opts ...grpc.CallOption) (*ID, error)
	SampleTesting(ctx context.Context, in *Command, opts ...grpc.CallOption) (*ID, error)
	// Return a list of all running (not reaped) commands by ID.
	Running(ctx context.Context, in *Empty, opts ...grpc.CallOption) (RCEAgent_RunningClient, error)
}

type rCEAgentClient struct {
	cc *grpc.ClientConn
}

func NewRCEAgentClient(cc *grpc.ClientConn) RCEAgentClient {
	return &rCEAgentClient{cc}
}

func (c *rCEAgentClient) Start(ctx context.Context, in *Command, opts ...grpc.CallOption) (*ID, error) {
	out := new(ID)
	err := c.cc.Invoke(ctx, "/mars.RCEAgent/Start", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rCEAgentClient) Wait(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/mars.RCEAgent/Wait", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rCEAgentClient) GetStatus(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/mars.RCEAgent/GetStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rCEAgentClient) Stop(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/mars.RCEAgent/Stop", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rCEAgentClient) GorTraffic(ctx context.Context, in *Command, opts ...grpc.CallOption) (*ID, error) {
	out := new(ID)
	err := c.cc.Invoke(ctx, "/mars.RCEAgent/GorTraffic", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rCEAgentClient) TrafficUpload(ctx context.Context, in *Command, opts ...grpc.CallOption) (*ID, error) {
	out := new(ID)
	err := c.cc.Invoke(ctx, "/mars.RCEAgent/TrafficUpload", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rCEAgentClient) Testing(ctx context.Context, in *Command, opts ...grpc.CallOption) (*ID, error) {
	out := new(ID)
	err := c.cc.Invoke(ctx, "/mars.RCEAgent/Testing", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rCEAgentClient) SampleTesting(ctx context.Context, in *Command, opts ...grpc.CallOption) (*ID, error) {
	out := new(ID)
	err := c.cc.Invoke(ctx, "/mars.RCEAgent/SampleTesting", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rCEAgentClient) Running(ctx context.Context, in *Empty, opts ...grpc.CallOption) (RCEAgent_RunningClient, error) {
	stream, err := c.cc.NewStream(ctx, &_RCEAgent_serviceDesc.Streams[0], "/mars.RCEAgent/Running", opts...)
	if err != nil {
		return nil, err
	}
	x := &rCEAgentRunningClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type RCEAgent_RunningClient interface {
	Recv() (*ID, error)
	grpc.ClientStream
}

type rCEAgentRunningClient struct {
	grpc.ClientStream
}

func (x *rCEAgentRunningClient) Recv() (*ID, error) {
	m := new(ID)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// RCEAgentServer is the server API for RCEAgent service.
type RCEAgentServer interface {
	// Start a command and immediately return its ID. Be sure to call Wait or Stop
	// to reap the command, else the agent will effectively leak memory by holding
	// unreaped commands. A command is considered running until reaped.
	Start(context.Context, *Command) (*ID, error)
	// Wait for a command to complete or be stopped, reap it, and return its final status.
	Wait(context.Context, *ID) (*Status, error)
	// Get the status of a command if it hasn't been reaped by calling Wait or Stop.
	GetStatus(context.Context, *ID) (*Status, error)
	// Stop then reap a command by sending it a SIGTERM signal.
	Stop(context.Context, *ID) (*Empty, error)
	// start gor task
	GorTraffic(context.Context, *Command) (*ID, error)
	// upload traffic to oss
	TrafficUpload(context.Context, *Command) (*ID, error)
	// start jmeter task
	Testing(context.Context, *Command) (*ID, error)
	SampleTesting(context.Context, *Command) (*ID, error)
	// Return a list of all running (not reaped) commands by ID.
	Running(*Empty, RCEAgent_RunningServer) error
}

func RegisterRCEAgentServer(s *grpc.Server, srv RCEAgentServer) {
	s.RegisterService(&_RCEAgent_serviceDesc, srv)
}

func _RCEAgent_Start_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Command)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RCEAgentServer).Start(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mars.RCEAgent/Start",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RCEAgentServer).Start(ctx, req.(*Command))
	}
	return interceptor(ctx, in, info, handler)
}

func _RCEAgent_Wait_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RCEAgentServer).Wait(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mars.RCEAgent/Wait",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RCEAgentServer).Wait(ctx, req.(*ID))
	}
	return interceptor(ctx, in, info, handler)
}

func _RCEAgent_GetStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RCEAgentServer).GetStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mars.RCEAgent/GetStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RCEAgentServer).GetStatus(ctx, req.(*ID))
	}
	return interceptor(ctx, in, info, handler)
}

func _RCEAgent_Stop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RCEAgentServer).Stop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mars.RCEAgent/Stop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RCEAgentServer).Stop(ctx, req.(*ID))
	}
	return interceptor(ctx, in, info, handler)
}

func _RCEAgent_GorTraffic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Command)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RCEAgentServer).GorTraffic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mars.RCEAgent/GorTraffic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RCEAgentServer).GorTraffic(ctx, req.(*Command))
	}
	return interceptor(ctx, in, info, handler)
}

func _RCEAgent_TrafficUpload_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Command)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RCEAgentServer).TrafficUpload(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mars.RCEAgent/TrafficUpload",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RCEAgentServer).TrafficUpload(ctx, req.(*Command))
	}
	return interceptor(ctx, in, info, handler)
}

func _RCEAgent_Testing_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Command)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RCEAgentServer).Testing(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mars.RCEAgent/Testing",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RCEAgentServer).Testing(ctx, req.(*Command))
	}
	return interceptor(ctx, in, info, handler)
}

func _RCEAgent_SampleTesting_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Command)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RCEAgentServer).SampleTesting(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mars.RCEAgent/SampleTesting",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RCEAgentServer).SampleTesting(ctx, req.(*Command))
	}
	return interceptor(ctx, in, info, handler)
}

func _RCEAgent_Running_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RCEAgentServer).Running(m, &rCEAgentRunningServer{stream})
}

type RCEAgent_RunningServer interface {
	Send(*ID) error
	grpc.ServerStream
}

type rCEAgentRunningServer struct {
	grpc.ServerStream
}

func (x *rCEAgentRunningServer) Send(m *ID) error {
	return x.ServerStream.SendMsg(m)
}

var _RCEAgent_serviceDesc = grpc.ServiceDesc{
	ServiceName: "mars.RCEAgent",
	HandlerType: (*RCEAgentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Start",
			Handler:    _RCEAgent_Start_Handler,
		},
		{
			MethodName: "Wait",
			Handler:    _RCEAgent_Wait_Handler,
		},
		{
			MethodName: "GetStatus",
			Handler:    _RCEAgent_GetStatus_Handler,
		},
		{
			MethodName: "Stop",
			Handler:    _RCEAgent_Stop_Handler,
		},
		{
			MethodName: "GorTraffic",
			Handler:    _RCEAgent_GorTraffic_Handler,
		},
		{
			MethodName: "TrafficUpload",
			Handler:    _RCEAgent_TrafficUpload_Handler,
		},
		{
			MethodName: "Testing",
			Handler:    _RCEAgent_Testing_Handler,
		},
		{
			MethodName: "SampleTesting",
			Handler:    _RCEAgent_SampleTesting_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Running",
			Handler:       _RCEAgent_Running_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "mars.proto",
}
