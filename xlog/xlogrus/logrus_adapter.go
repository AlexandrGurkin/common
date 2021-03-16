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

func NewXLogrus(cfg xlog.LoggerCfg) xlog.Logger {
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
	//logrus.AddHook(hooks.LogProm{})
	return &xrus{*logrus.NewEntry(logrus.StandardLogger())}
}

func (x *xrus) Debug(msg string) {
	x.Debugf("%s", msg)
}

func (x *xrus) Info(msg string) {
	x.Infof("%s", msg)
}

func (x *xrus) Warn(msg string) {
	x.Warnf("%s", msg)
}
func (x *xrus) Trace(msg string) {
	x.Tracef("%s", msg)
}

func (x *xrus) Error(msg string) {
	x.Errorf("%s", msg)
}

func (x *xrus) Fatal(msg string) {
	x.Fatalf("%s", msg)
}
