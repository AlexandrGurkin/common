package xlogrus

import (
	"time"

	"github.com/AlexandrGurkin/common/xlog"
	"github.com/sirupsen/logrus"
)

type xrus struct {
	logrus.Entry
}

func (x *xrus) WithXFields(fields xlog.Fields) xlog.Logger {
	return &xrus{*x.WithFields(logrus.Fields(fields))}
}

func (x *xrus) WithXField(key string, value interface{}) xlog.Logger {
	return &xrus{*x.WithField(key, value)}
}

func NewXLogrus(cfg xlog.LoggerCfg) (xlog.Logger, error) {
	lvl, err := logrus.ParseLevel(cfg.Level)
	if err != nil {
		logrus.SetLevel(logrus.DebugLevel)
	} else {
		logrus.SetLevel(lvl)
	}
	logrus.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: time.RFC3339Nano,
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyMsg: "message",
		},
	})
	logrus.SetOutput(cfg.Out)

	return &xrus{*logrus.NewEntry(logrus.StandardLogger())}, nil
}

func (x *xrus) Debug(msg string) {
	x.Entry.Debug(msg)
}

func (x *xrus) Info(msg string) {
	x.Entry.Info(msg)
}

func (x *xrus) Warn(msg string) {
	x.Entry.Warn(msg)
}
func (x *xrus) Trace(msg string) {
	x.Entry.Trace(msg)
}

func (x *xrus) Error(msg string) {
	x.Entry.Error(msg)
}

func (x *xrus) Fatal(msg string) {
	x.Entry.Fatal(msg)
}
