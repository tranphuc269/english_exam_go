package app_logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	// Logger instance
	Logger *zap.Logger
)

// Init Zap-Logger initialization process
func Init() {
	e, err := getEnv()
	if err != nil {
		panic("Can not init Logger. error:" + err.Error())
	}

	var outputPaths []string
	if e.filePath != "" {
		outputPaths = append(outputPaths, e.filePath)
	}

	if e.stdout {
		outputPaths = append(outputPaths, "stdout")
	}

	c := zap.Config{
		OutputPaths: outputPaths,
		Level:       zap.NewAtomicLevelAt(e.level),
		Encoding:    "json",
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:        "Time",
			LevelKey:       "Level",
			NameKey:        "Name",
			CallerKey:      "Caller",
			MessageKey:     "Message",
			StacktraceKey:  "St",
			EncodeLevel:    zapcore.CapitalLevelEncoder,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeDuration: zapcore.StringDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
	}

	if Logger, err = c.Build(); err != nil {
		panic("Can not init Logger. error:" + err.Error())
	}
}
