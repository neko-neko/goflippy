package log

import (
	"os"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

var logger log.Logger

func init() {
	l := log.NewLogfmtLogger(os.Stdout)
	if os.Getenv("DEBUG") == "true" {
		l = level.NewFilter(l, level.AllowAll())
	} else {
		l = level.NewFilter(l, level.AllowInfo())
	}
	logger = log.With(l, "timestamp", log.DefaultTimestamp, "caller", log.Caller(4))
}

// Debug outputs log for debug
func Debug(keyvals ...interface{}) {
	level.Debug(logger).Log(keyvals...)
}

// DebugWithMsg outputs debug level log with message
func DebugWithMsg(msg string, keyvals ...interface{}) {
	Debug(append(keyvals, "message", msg)...)
}

// Info outputs log for info
func Info(keyvals ...interface{}) {
	level.Info(logger).Log(keyvals...)
}

// InfoWithMsg outputs info level log with message
func InfoWithMsg(msg string, keyvals ...interface{}) {
	Info(append(keyvals, "message", msg)...)
}

// Warn outputs log for warn
func Warn(keyvals ...interface{}) {
	level.Warn(logger).Log(keyvals...)
}

// WarnWithMsg outputs warn level log with message
func WarnWithMsg(msg string, keyvals ...interface{}) {
	Warn(append(keyvals, "message", msg)...)
}

// Error outputs log for error
func Error(keyvals ...interface{}) {
	level.Error(logger).Log(keyvals...)
}

// ErrorWithErr outputs log for error
func ErrorWithErr(err error, keyvals ...interface{}) {
	Error(append(keyvals, "err", err)...)
}
