package app_logger

import (
	"os"
	"strconv"

	"go.uber.org/zap/zapcore"
)

type envVals struct {
	filePath string
	stdout   bool
	level    zapcore.Level
}

// getEnv Loggerに関する設定を環境変数から取得
func getEnv() (*envVals, error) {
	res := new(envVals)

	// config path for log
	res.filePath = os.Getenv("LOGGER_FILE_PATH")

	// Standard output setting
	var err error
	res.stdout, err = strconv.ParseBool(os.Getenv("LOGGER_STDOUT"))
	if err != nil {
		res.stdout = true
	}

	// Log level
	level := os.Getenv("LOGGER_LEVEL")

	switch level {
	case "debug":
		res.level = zapcore.DebugLevel
	case "error":
		res.level = zapcore.ErrorLevel
	default:
		res.level = zapcore.InfoLevel
	}

	return res, nil
}
