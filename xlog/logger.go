// Package xlog contain interface for Logger
//
package xlog

// Logger interface for integration_common
//
type Logger interface {
	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Tracef(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Debug(args ...interface{})
	Info(args ...interface{})
	Warn(args ...interface{})
	Trace(args ...interface{})
	Error(args ...interface{})

	WithFields(fields Fields) Logger
	WithField(key string, value interface{}) Logger
}

type Fields map[string]interface{}
