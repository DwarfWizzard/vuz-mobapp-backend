package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func init() {
	lvl := os.Getenv("LOG_LEVEL")
	logger := newLogger(lvl)
	zap.ReplaceGlobals(logger)
}

func newLogger(level string) *zap.Logger {
	lvl, err := zapcore.ParseLevel(level)
	if err != nil {
		lvl = zap.InfoLevel
	}

	encConfig := zapcore.EncoderConfig{
		MessageKey:     "message",
		LevelKey:       "level",
		TimeKey:        "timestamp",
		NameKey:        "logger",
		CallerKey:      "origin",
		StacktraceKey:  zapcore.OmitKey,
		SkipLineEnding: false,
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.NanosDurationEncoder,
		EncodeName:     zapcore.FullNameEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	logger := zap.New(zapcore.NewCore(zapcore.NewJSONEncoder(encConfig), zapcore.Lock(os.Stdout), lvl))
	return logger
}

func Logger() *zap.Logger {
	return zap.L().Named("Logger")
}
