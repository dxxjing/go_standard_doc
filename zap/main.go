package main

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"time"
)


func main() {
	//test1()
	//test2()
	customize()
	zap.L().Info("test")
}



func customize(){
	hook := lumberjack.Logger{
		Filename:   "./logs/spikeProxy1.log", // 日志文件路径
		MaxSize:    128,                      // 每个日志文件保存的最大尺寸 单位：M
		MaxBackups: 30,                       // 日志文件最多保存多少个备份
		MaxAge:     7,                        // 文件最多保存多少天
		Compress:   true,                     // 是否压缩
	}

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "linenum",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,  // 小写编码器
		EncodeTime:     zapcore.ISO8601TimeEncoder,     // ISO8601 UTC 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder, //
		//EncodeCaller:   zapcore.FullCallerEncoder,      // 全路径编码器
		EncodeCaller:   MyCallerEncoder,      // 全路径编码器
		EncodeName:     zapcore.FullNameEncoder,
		ConsoleSeparator: "|",
	}
	//设置日志级别
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(zap.InfoLevel)

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),                                           // 编码器配置 json格式
		//zapcore.NewConsoleEncoder(encoderConfig),		//普通文本格式
		//zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook)), // 打印到控制台和文件
		zapcore.AddSync(&hook),	//打印到文件
		atomicLevel,                                                                     // 日志级别
	)
	// 开启开发模式，堆栈跟踪
	caller := zap.AddCaller()
	// 开启文件及行号
	development := zap.Development()
	// 设置初始化字段
	filed := zap.Fields(
		zap.String("serviceName", "hippo-server"),
		zap.Int("uid", 3688),
		)

	// 构造日志
	logger := zap.New(core, caller, development, filed)
	zap.ReplaceGlobals(logger)
	logger.Info("hello zap", zap.String("aaa", "bbb"))

}

func MyCallerEncoder(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendInt(caller.Line)
	enc.AppendString(caller.String())
}

//输出日志到控制台 格式化输出
func test1() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		fmt.Println(err)
		return
	}
	//刷新所有缓冲区的日志 程序退出前必须调用
	defer logger.Sync()

	logger.Info("无法获取网址",
		zap.String("url", "http://www.baidu.com"),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)
	//2021-01-07T18:09:41.231+0800    INFO    zap/main.go:22  无法获取网址    {"url": "http://www.baidu.com", "attempt": 3, "backoff": "1s"}
}

//输出日志到控制台 json序列化输出
func test2(){
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	logger.Info("无法获取网址",
		zap.String("url", "http://www.baidu.com"),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)
	//{"level":"info","ts":1610014372.893056,"caller":"zap/main.go:36","msg":"无法获取网址","url":"http://www.baidu.com","attempt":3,"backoff":1}
}
