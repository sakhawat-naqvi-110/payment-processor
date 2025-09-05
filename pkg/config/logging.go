package config

import (
	"context"
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	logger2 "gorm.io/gorm/logger"
)

var logger *zap.Logger

// InitializeLogger sets up a global logger using the zap library.
func InitializeLogger() {

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	fileEncoder := zapcore.NewConsoleEncoder(encoderConfig)
	// Use a JSON encoder for both file and console
	jsonEncoder := zapcore.NewJSONEncoder(encoderConfig)
	logFile, _ := os.OpenFile("text.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	writer := zapcore.AddSync(logFile)
	defaultLogLevel := zapcore.DebugLevel
	core := zapcore.NewTee(
		zapcore.NewCore(fileEncoder, writer, defaultLogLevel),
		zapcore.NewCore(jsonEncoder, zapcore.AddSync(os.Stdout), defaultLogLevel),
	)
	logger = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
	zap.ReplaceGlobals(logger)
}

// GetLogger returns the globally initialized zap.Logger instance.
func GetLogger() *zap.Logger {
	return logger
}

// ZapLogger is a custom logger that implements GORM's logger2.Interface.
type ZapLogger struct{}

// LogMode method sets the logging mode, returning the same logger for GORM compatibility.
func (l ZapLogger) LogMode(level logger2.LogLevel) logger2.Interface {
	return l
}

// Info Logs informational messages with optional context and additional data.
func (l ZapLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	logger.Info(msg, zap.Any("context", ctx), zap.Any("data", data))
}

// Warn Logs warning messages with optional context and additional data.
func (l ZapLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	logger.Warn(msg, zap.Any("context", ctx), zap.Any("data", data))
}

// Error Logs error messages with optional context and additional data.
func (l ZapLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	logger.Error(msg, zap.Any("context", ctx), zap.Any("data", data))
}

// Trace Logs detailed SQL query information and handles logging errors related to SQL execution.
func (l ZapLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {

	sql, rows := fc()
	if err != nil {
		logger.Error("Query error", zap.Error(err), zap.String("sql", sql), zap.Int64("rows", rows))
	} else {
		logger.Info("Query OK", zap.String("sql", sql), zap.Int64("rows", rows))
	}
}
