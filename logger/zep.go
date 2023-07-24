package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

type TestLogger struct {
	logger *zap.Logger
}

var Log TestLogger

//	func (l *TestLogger) FileLogger(filename string) *zap.Logger {
//		config := zap.NewProductionEncoderConfig()
//
//		config.EncodeTime = zapcore.ISO8601TimeEncoder
//		// fileLog 남기기 (JSON 형식)
//		fileEncoder := zapcore.NewJSONEncoder(config)
//		// console 남기기
//		consoleEncoder := zapcore.NewConsoleEncoder(config)
//		logFile, _ := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
//		writer := zapcore.AddSync(logFile)
//
//		//logLevel 설정
//		defaultLogLevel := zapcore.InfoLevel
//
//		// 두가지 core를 사용시 사용하는 NewTee
//		core := zapcore.NewTee(
//			zapcore.NewCore(fileEncoder, writer, defaultLogLevel),
//			zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), defaultLogLevel),
//		)
//
//		logger := zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
//
//		return logger
//	}
func NewLogger() *zap.SugaredLogger {
	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.TimeKey = "timestamp"
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder

	config := zap.Config{
		Level:             zap.NewAtomicLevelAt(zap.InfoLevel),
		Development:       false,
		DisableCaller:     false,
		DisableStacktrace: false,
		Sampling:          nil,
		Encoding:          "console",
		EncoderConfig:     encoderCfg,
		OutputPaths: []string{
			"stderr",
		},
		ErrorOutputPaths: []string{
			"stderr",
		},
		InitialFields: map[string]interface{}{
			"pid": os.Getpid(),
		},
	}

	return zap.Must(config.Build()).Sugar()
}
