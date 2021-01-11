package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"time"
)

func main(){
	//"%datetime%|%channel%|%level_name%|%trace_id%|%version%|%uid%|%url%|%params%|%file%|%class%|%line%|%message%|%data%\n";
	encoderConfig := zapcore.EncoderConfig{
		MessageKey:  "msg",
		LevelKey:    "level",
		CallerKey:    "file",
		TimeKey:     "time",
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format("2006-01-02 15:04:05"))
		},
		EncodeLevel: zapcore.LowercaseLevelEncoder,
		EncodeCaller: func(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(caller.File)
			enc.AppendInt(caller.Line)

		},
		EncodeDuration: func(d time.Duration, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendInt64(int64(d) / 1000000)
		},
		ConsoleSeparator: "|",
	}
	// 实现两个判断日志等级的interface
	infoLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl < zapcore.ErrorLevel
	})

	warnLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl == zapcore.WarnLevel
	})

	errLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.ErrorLevel
	})
	// 设置初始化字段
	initFields := zap.Fields(
		zap.String("serviceName", "hippo-server"),
		zap.Int("uid", 3688),
	)
	// 获取 info、error日志文件的io.Writer
	infoWriter := getWriter("info")
	errWriter := getWriter("error")
	warnWriter := getWriter("warning")
	core := zapcore.NewTee(
		// 打印在文件中
		zapcore.NewCore(zapcore.NewConsoleEncoder(encoderConfig), zapcore.AddSync(infoWriter), infoLevel),
		zapcore.NewCore(zapcore.NewConsoleEncoder(encoderConfig), zapcore.AddSync(warnWriter), warnLevel),
		zapcore.NewCore(zapcore.NewConsoleEncoder(encoderConfig), zapcore.AddSync(errWriter), errLevel),
	)
	logger := zap.New(core).WithOptions(zap.AddCaller(),initFields)
	logger.Info("test info")
}

func getWriter(level string) *lumberjack.Logger {
	filename := "./logs/test-" +  level +".log"
	hook := lumberjack.Logger{
		Filename:   filename, // ⽇志⽂件路径
		MaxSize:    1024,    // megabytes
		MaxBackups: 3,       // 最多保留3个备份
		MaxAge:     7,       //days
		Compress:   true,    // 是否压缩 disabled by default
	}

	return &hook

}