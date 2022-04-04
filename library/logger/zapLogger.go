package logger

import (
	agentConfig "github.com/lehoon/devops-agent/config"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	agentLogger *zap.SugaredLogger
)

func init() {
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "log",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "trace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	atomLevel := zap.NewAtomicLevelAt(zapcore.Level(agentConfig.GetLoggerLevel()))

	loggerHook := lumberjack.Logger{
		Filename: agentConfig.GetLoggerPath(),
		MaxSize: agentConfig.GetLoggerMaxSize(),
		MaxAge: agentConfig.GetLoggerMaxAge(),
		MaxBackups: agentConfig.GetLoggerMaxBackup(),
		Compress: agentConfig.GetLoggerCompress(),
	}

	loggerCore := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(&loggerHook)),
		atomLevel,
	)

	caller := zap.AddCaller()
	development := zap.Development()
	field := zap.Fields(zap.String("serviceName", agentConfig.GetLoggerServiceName()))
	logger := zap.New(loggerCore, caller, development, field)
	logger.Info("logger初始化完成")
	agentLogger = logger.Sugar()

}

func Log() *zap.SugaredLogger {
	return agentLogger
}

func Sync() {
	agentLogger.Sync()
}
