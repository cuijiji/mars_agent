package code

/**
error code
*/
const (
	ErrorFlag          int = 1
	ConsulConfigIsNull int = 2
	GrpcStartError     int = 3
	GrpcRegisterError  int = 4
)

var codeEnumMap = map[int]string{
	ErrorFlag:          "启动指令错误",
	ConsulConfigIsNull: "consul配置为nil",
	GrpcStartError:     "grpc启动错误",
	GrpcRegisterError:  "grpc consul 注册失败",
}

func GetErrorMsg(key int) string {
	return codeEnumMap[key]
}
