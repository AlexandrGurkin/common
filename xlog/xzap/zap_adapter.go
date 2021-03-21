package xzap

import (
	"time"

	"github.com/AlexandrGurkin/common/xlog"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type xZap struct {
	l *zap.Logger
	s *zap.SugaredLogger
}

func NewXZap(cfg xlog.LoggerCfg) (xlog.Logger, error) {
	if cfg.TimeFormat == "" {
		cfg.TimeFormat = time.RFC3339Nano
	}
	encoderCfg := zapcore.EncoderConfig{
		MessageKey:     "message",
		LevelKey:       "level",
		NameKey:        "logger",
		TimeKey:        "time",
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.TimeEncoderOfLayout(cfg.TimeFormat),
		EncodeDuration: zapcore.StringDurationEncoder,
	}

	level := zapcore.DebugLevel
	if cfg.Level != "" {
		if cfg.Level == "trace" { //zap(
			cfg.Level = "debug"
		}
		if err := level.Set(cfg.Level); err != nil {
			return nil, xlog.ErrorInitLogger.Wrap(err)
		}
	}

	var out zapcore.WriteSyncer
	var ok bool
	if out, ok = cfg.Out.(zapcore.WriteSyncer); !ok {
		out = zapcore.AddSync(cfg.Out)
	}

	core := zapcore.NewCore(zapcore.NewJSONEncoder(encoderCfg), out, level)
	logger := zap.New(core)
	sugar := logger.Sugar()

	return &xZap{logger, sugar}, nil
}

func (z *xZap) Debugf(format string, args ...interface{}) {
	z.s.Debugf(format, args...)
}

func (z *xZap) Infof(format string, args ...interface{}) {
	z.s.Infof(format, args...)
}

func (z *xZap) Warnf(format string, args ...interface{}) {
	z.s.Warnf(format, args...)
}

func (z *xZap) Tracef(format string, args ...interface{}) {
	z.s.Debugf(format, args...)
}

func (z *xZap) Errorf(format string, args ...interface{}) {
	z.s.Errorf(format, args...)
}

func (z *xZap) Fatalf(format string, args ...interface{}) {
	z.s.Fatalf(format, args...)
}

func (z *xZap) Debug(msg string) {
	z.l.Debug(msg)
}

func (z *xZap) Info(msg string) {
	z.l.Info(msg)
}

func (z *xZap) Warn(msg string) {
	z.l.Warn(msg)
}

func (z *xZap) Trace(msg string) {
	z.l.Debug(msg)
}

func (z *xZap) Error(msg string) {
	z.l.Error(msg)
}

func (z *xZap) Fatal(msg string) {
	z.l.Fatal(msg)
}

func (z *xZap) WithXFields(fields xlog.Fields) xlog.Logger {
	fs := make([]interface{}, 0, len(fields))
	for k, v := range fields {
		fs = append(fs, k)
		fs = append(fs, v)
	}
	sugar := z.s.With(fs...)
	log := sugar.Desugar()
	return &xZap{
		l: log,
		s: sugar,
	}
}

func (z *xZap) WithXField(key string, value interface{}) xlog.Logger {
	sugar := z.s.With(key, value)
	log := sugar.Desugar()
	return &xZap{
		l: log,
		s: sugar,
	}
}
