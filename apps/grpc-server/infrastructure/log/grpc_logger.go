package log

import (
	"fmt"
	"strings"

	"go.uber.org/zap"
	"google.golang.org/grpc/grpclog"
)

// workaround
var skipMessages = []string{
	"transport is closing",
	"connection reset by peer",
}

// ReplaceGrpcLogger ReplaceGrpcLogger
func ReplaceGrpcLogger(logger *zap.Logger) {
	zgl := &zapGrpcLogger{logger.With(zap.String("system", "grpc"), zap.Bool("grpc_log", true))}
	grpclog.SetLogger(zgl)
}

type zapGrpcLogger struct {
	logger *zap.Logger
}

func (l *zapGrpcLogger) Fatal(args ...interface{}) {
	l.logger.Fatal(fmt.Sprint(args...))
}

func (l *zapGrpcLogger) Fatalf(format string, args ...interface{}) {
	l.logger.Fatal(fmt.Sprintf(format, args...))
}

func (l *zapGrpcLogger) Fatalln(args ...interface{}) {
	l.logger.Fatal(fmt.Sprint(args...))
}

func (l *zapGrpcLogger) Print(args ...interface{}) {
	msg := fmt.Sprint(args...)
	if skip(msg) {
		return
	}
	l.logger.Info(msg)
}

func (l *zapGrpcLogger) Printf(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	if skip(msg) {
		return
	}
	l.logger.Info(msg)

}

func (l *zapGrpcLogger) Println(args ...interface{}) {
	msg := fmt.Sprint(args...)
	if skip(msg) {
		return
	}
	l.logger.Info(msg)
}

func skip(msg string) bool {
	for i := 0; i < len(skipMessages); i++ {
		if !strings.Contains(msg, skipMessages[i]) {
			return true
		}
	}
	return false
}
