package log

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/middleware"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// LogFormatter LogFormatter
type logFormatter struct {
	logger *zap.Logger
}

// NewLogFormatter NewLogFormatter
func NewLogFormatter(l *zap.Logger) middleware.LogFormatter {
	return &logFormatter{l}
}

// NewLogEntry LogEntry
func (f *logFormatter) NewLogEntry(r *http.Request) middleware.LogEntry {
	// TODO: fix hardcode
	if r.RequestURI == "/health" {
		return &logSkipEntry{}
	}

	entry := &logEntry{f.logger}

	scheme := "http"
	if r.TLS != nil {
		scheme = "https"
	}

	fields := []zapcore.Field{
		zap.String("ts", time.Now().UTC().Format(time.RFC1123)),
		zap.String("http_scheme", scheme),
		zap.String("http_proto", r.Proto),
		zap.String("http_method", r.Method),
		zap.String("remote_addr", r.RemoteAddr),
		zap.String("user_agent", r.UserAgent()),
		zap.Int64("content_length", r.ContentLength),
		zap.String("uri", fmt.Sprintf("%s://%s%s", scheme, r.Host, r.RequestURI)),
	}
	entry.logger = entry.logger.With(fields...)
	return entry
}

type logEntry struct {
	logger *zap.Logger
}

func (l *logEntry) Write(status, bytes int, elapsed time.Duration) {
	l.logger.
		With(zap.Float64("response_time_ms", float64(elapsed.Nanoseconds())/1000000.0)).
		With(zap.Int("status_code", status)).
		With(zap.Int("bytes_out", bytes)).
		Info("logging http request")
}

func (l *logEntry) Panic(v interface{}, stack []byte) {
	// TODO: test
	l.logger.With([]zapcore.Field{
		zap.String("stack", string(stack)),
		zap.String("panic", fmt.Sprintf("%+v", v)),
	}...).Panic("panic")
}

type logSkipEntry struct{}

func (l *logSkipEntry) Write(status, bytes int, elapsed time.Duration) {}
func (l *logSkipEntry) Panic(v interface{}, stack []byte)              {}
