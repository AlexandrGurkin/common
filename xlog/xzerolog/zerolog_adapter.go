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
	if cfg.TimeFormat == "" {
		cfg.TimeFormat = time.RFC3339Nano
	}
	zerolog.TimeFieldFormat = cfg.TimeFormat
	logger := zerolog.New(cfg.Out).With().Timestamp().Logger()
	lvl, err := zerolog.ParseLevel(cfg.Level)
	if err != nil {
		return nil, xlog.ErrorInitLogger.Wrap(err)
	}
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
	switch v := value.(type) {
	case int8:
		return &zLog{z.l.With().Int8(key, v).Logger()}
	case int:
		return &zLog{z.l.With().Int(key, v).Logger()}
	case int16:
		return &zLog{z.l.With().Int16(key, v).Logger()}
	case int32:
		return &zLog{z.l.With().Int32(key, v).Logger()}
	case int64:
		return &zLog{z.l.With().Int64(key, v).Logger()}
	case uint8:
		return &zLog{z.l.With().Uint8(key, v).Logger()}
	case uint:
		return &zLog{z.l.With().Uint(key, v).Logger()}
	case uint16:
		return &zLog{z.l.With().Uint16(key, v).Logger()}
	case uint32:
		return &zLog{z.l.With().Uint32(key, v).Logger()}
	case uint64:
		return &zLog{z.l.With().Uint64(key, v).Logger()}
	case string:
		return &zLog{z.l.With().Str(key, v).Logger()}
	case float32:
		return &zLog{z.l.With().Float32(key, v).Logger()}
	case float64:
		return &zLog{z.l.With().Float64(key, v).Logger()}
	case bool:
		return &zLog{z.l.With().Bool(key, v).Logger()}
	default:
		f := make(map[string]interface{})
		f[key] = value
		return &zLog{z.l.With().Fields(f).Logger()}
	}

}
