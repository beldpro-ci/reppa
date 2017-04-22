package common

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var logger *zap.Logger

func init() {
	var l *zap.Logger
	var err error

	if os.Getenv("DEBUG") != "" {
		config := zap.NewDevelopmentConfig()
		config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		l, err = config.Build()
	} else {
		l, err = zap.NewProduction()
	}
	if err != nil {
		panic(err)
	}

	logger = l
}

func GetLogger() *zap.Logger {
	return logger
}
