package log

import (
	"context"
	"fmt"
	"sync"

	"go.uber.org/zap/zapcore"

	"go.uber.org/zap"
)

// Logger Logger
var (
	Logger *zap.Logger
	once   = new(sync.Once)
)

// Setup Setup
func Setup() {
	once.Do(func() {
		// TODO: load onfig/log.json
		config := zap.NewProductionConfig()
		config.OutputPaths = []string{"stdout"}
		l, _ := config.Build(
			zap.AddCaller(),
			zap.AddCallerSkip(2),
		)
		Logger = l
	})
}

// Debug Debug
func Debug(ctx context.Context, msg string) {
	log(ctx, zap.DebugLevel, msg)
}

// Debugf Debugf
func Debugf(ctx context.Context, format string, a ...interface{}) {
	log(ctx, zap.DebugLevel, fmt.Sprintf(format, a...))
}

// Info Info
func Info(ctx context.Context, msg string) {
	log(ctx, zap.InfoLevel, msg)
}

// Infof Infof
func Infof(ctx context.Context, format string, a ...interface{}) {
	log(ctx, zap.InfoLevel, fmt.Sprintf(format, a...))
}

// Warn Warn
func Warn(ctx context.Context, msg string) {
	log(ctx, zap.WarnLevel, msg)
}

// Warnf Warnf
func Warnf(ctx context.Context, format string, a ...interface{}) {
	log(ctx, zap.WarnLevel, fmt.Sprintf(format, a...))
}

// Error Error
func Error(ctx context.Context, msg string) {
	log(ctx, zap.ErrorLevel, msg)
}

// Errorf Errorf
func Errorf(ctx context.Context, format string, a ...interface{}) {
	log(ctx, zap.ErrorLevel, fmt.Sprintf(format, a...))
}

// Fatal Fatal
func Fatal(ctx context.Context, msg string) {
	log(ctx, zap.FatalLevel, msg)
}

// Fatalf Fatalf
func Fatalf(ctx context.Context, format string, a ...interface{}) {
	log(ctx, zap.FatalLevel, fmt.Sprintf(format, a...))
}

// log
func log(ctx context.Context, level zapcore.Level, msg string) {

	// TODO: common format
	switch level {
	case zap.DebugLevel:
		Logger.Debug(msg)
	case zap.InfoLevel:
		Logger.Info(msg)
	case zap.WarnLevel:
		Logger.Warn(msg)
	case zap.ErrorLevel:
		Logger.Error(msg)
	case zap.FatalLevel:
		Logger.Fatal(msg)
	}
}
