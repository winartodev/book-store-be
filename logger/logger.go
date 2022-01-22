package logger

import (
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	logger *zap.Logger
	once   sync.Once
)

// Fileds holds key and value to be written to log
type Fields map[string]interface{}

// init is a method to initialize logger instance
func Init() {
	once.Do(func() {
		logger, _ = zap.NewProduction()
	})
}

// Info writes log with severity level info
func Info(message string, fields Fields) {
	zapfileds := []zapcore.Field{}
	for k, v := range fields {
		zapfileds = append(zapfileds, zap.Any(k, v))
	}
	logger.Info(message, zapfileds...)
}

// Error writes log with severity level error
func Error(err error, fileds Fields) {
	zapfileds := []zapcore.Field{}
	for k, v := range fileds {
		zapfileds = append(zapfileds, zap.Any(k, v))
	}
	logger.Error(err.Error(), zapfileds...)
}
