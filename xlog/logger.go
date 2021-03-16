// Package xlog contain interface for Logger
//
package xlog

import (
	"io"
	"sync/atomic"
)

// Logger interface for integration_common
//
type Logger interface {
	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Tracef(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
	Debug(msg string)
	Info(msg string)
	Warn(msg string)
	Trace(msg string)
	Error(msg string)
	Fatal(msg string)

	WithXFields(fields Fields) Logger
	WithXField(key string, value interface{}) Logger
}

type Fields map[string]interface{}

type LoggerCfg struct {
	Level string
	Out   io.Writer
}

type BlackholeStream struct {
	writeCount uint64
}

func (s *BlackholeStream) WriteCount() uint64 {
	return atomic.LoadUint64(&s.writeCount)
}

func (s *BlackholeStream) Write(p []byte) (int, error) {
	atomic.AddUint64(&s.writeCount, 1)
	return len(p), nil
}
