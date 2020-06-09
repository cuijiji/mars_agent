package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"strings"
)

var (
	Logger      *zap.Logger
	HostName, _ = os.Hostname()
)

func InitLogConfig(logPath string, logName string) {
	// 去除log目录
	i := strings.LastIndex(logPath, "/")
	if i == len(logPath) {
		logPath = strings.TrimRight(logPath, "/")
	}
	hook := lumberjack.Logger{
		Filename:   logPath + "/" + logName + ".log", // 日志文件路径
		MaxSize:    1000,                             // 每个日志文件保存的最大尺寸 单位：M
		MaxBackups: 30,                               // 日志文件最多保存多少个备份
		MaxAge:     7,                                // 文件最多保存多少天
		Compress:   true,                             // 是否压缩
	}

	fileEncoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "file",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,  // 小写编码器
		EncodeTime:     zapcore.ISO8601TimeEncoder,     // ISO8601 UTC 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder, //
		EncodeCaller:   zapcore.ShortCallerEncoder,     // 全路径编码器
		EncodeName:     zapcore.FullNameEncoder,
	}

	pe := zap.NewProductionEncoderConfig()
	pe.EncodeTime = zapcore.ISO8601TimeEncoder
	pe.EncodeLevel = zapcore.LowercaseColorLevelEncoder
	consoleEncoder := zapcore.NewConsoleEncoder(pe)

	// 设置日志级别
	//atomicLevel := zap.NewAtomicLevel()
	//atomicLevel.SetLevel(zap.InfoLevel)

	core := zapcore.NewTee(
		zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), zap.InfoLevel),
		zapcore.NewCore(zapcore.NewJSONEncoder(fileEncoderConfig), zapcore.AddSync(&hook), zap.InfoLevel),
	)

	//core := zapcore.NewCore(
	//	zapcore.NewJSONEncoder(fileEncoderConfig),                                       // 编码器配置
	//	zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook)), // 打印到控制台和文件
	//	atomicLevel,                                                                     // 日志级别
	//)

	// 开启开发模式，堆栈跟踪
	caller := zap.AddCaller()
	// 开启文件及行号
	development := zap.Development()
	// 设置初始化字段
	filed := zap.Fields(zap.String("service_name", logName), zap.String("hn", HostName))
	// 构造日志
	logger := zap.New(core, caller, development, filed)

	logger.Info("log 初始化成功")
	//logger.Sugar().Info("log 初始化成功")
	//logger.Info("无法获取网址",
	//	zap.String("url", "http://www.baidu.com"),
	//	zap.Int("attempt", 3),
	//	zap.Duration("backoff", time.Second))
	Logger = logger
	defer logger.Sync()
}
