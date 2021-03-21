package xzerolog

import (
	"time"

	"github.com/AlexandrGurkin/common/xlog"
	"github.com/rs/zerolog"
)

type zLog struct {
	l zerolog.Logger
}

func NewXZerolog(cfg xlog.LoggerCfg) (xlog.Logger, error) {
	zerolog.TimeFieldFormat = time.RFC3339Nano
	logger := zerolog.New(cfg.Out).With().Timestamp().Logger()
	lvl, _ := zerolog.ParseLevel(cfg.Level)
	zerolog.SetGlobalLevel(lvl)
	return &zLog{logger}, nil
}

func (z *zLog) Debugf(format string, args ...interface{}) {
	z.l.Debug().Msgf(format, args...)
}

func (z *zLog) Infof(format string, args ...interface{}) {
	z.l.Info().Msgf(format, args...)
}

func (z *zLog) Warnf(format string, args ...interface{}) {
	z.l.Warn().Msgf(format, args...)
}

func (z *zLog) Tracef(format string, args ...interface{}) {
	z.l.Trace().Msgf(format, args...)
}

func (z *zLog) Errorf(format string, args ...interface{}) {
	z.l.Error().Msgf(format, args...)
}

func (z *zLog) Fatalf(format string, args ...interface{}) {
	z.l.Fatal().Msgf(format, args...)
}

func (z *zLog) Debug(msg string) {
	z.l.Debug().Msg(msg)
}

func (z *zLog) Info(msg string) {
	z.l.Info().Msg(msg)
}

func (z *zLog) Warn(msg string) {
	z.l.Warn().Msg(msg)
}

func (z *zLog) Trace(msg string) {
	z.l.Trace().Msg(msg)
}

func (z *zLog) Error(msg string) {
	z.l.Error().Msg(msg)
}

func (z *zLog) Fatal(msg string) {
	z.l.Fatal().Msg(msg)
}

func (z *zLog) WithXFields(fields xlog.Fields) xlog.Logger {
	return &zLog{z.l.With().Fields(fields).Logger()}
}

func (z *zLog) WithXField(key string, value interface{}) xlog.Logger {
	f := make(map[string]interface{})
	f[key] = value
	return &zLog{z.l.With().Fields(f).Logger()}
}
