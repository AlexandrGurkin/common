package xzap

import (
	"fmt"
	"time"

	"github.com/AlexandrGurkin/common/xlog"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type xZap struct {
	l *zap.Logger
	s *zap.SugaredLogger
}

func NewLoggerWithSample(logLevel string, debug bool, first, thereafter int, samplerDuration time.Duration) (*zap.Logger, error) {
	var (
		prodEncoder    zapcore.Encoder
		consoleEncoder zapcore.Encoder
	)

	if debug {
		consoleEncoder = zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
	} else {
		prodEncoder = zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
	}

	cz, _, err := zap.Open("stdout")
	if err != nil {
		return nil, err
	}

	level := zapcore.DebugLevel
	if logLevel != "" {
		if err := level.Set(logLevel); err != nil {
			fmt.Printf("wrong log level[%s]. valid are: [DEBUG,INFO,WARN,ERROR,DPANIC,PANIC,FATAL]", logLevel)
			return nil, err
		}
	}

	consolePriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= level
	})

	var core zapcore.Core
	if debug {
		core = zapcore.NewCore(consoleEncoder, cz, consolePriority)
	} else {
		core = zapcore.NewCore(prodEncoder, cz, consolePriority)
	}

	queue := make(chan string, 2)
	for elem := range queue {
		fmt.Println(elem)
	}

	sampler := zapcore.NewSamplerWithOptions(
		core,
		samplerDuration,
		first,
		thereafter,
	)

	var logger *zap.Logger
	if debug {
		logger = zap.New(sampler, zap.Development())
	} else {
		logger = zap.New(sampler)
	}

	return logger, nil
}

func NewXZap(cfg xlog.LoggerCfg) (xlog.Logger, error) {
	//zerolog.TimeFieldFormat = time.RFC3339Nano
	encoderCfg := zapcore.EncoderConfig{
		MessageKey:     "message",
		LevelKey:       "level",
		NameKey:        "logger",
		TimeKey:        "time",
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.RFC3339NanoTimeEncoder,
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

	core := zapcore.NewCore(zapcore.NewJSONEncoder(encoderCfg), cfg.Out, level)
	//core = zapcore.NewSamplerWithOptions(core, 10 * time.Second, 1, 2)
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
