package middlewares

import "github.com/AlexandrGurkin/common/xlog"

type MiddlewareConfig struct {
	Logger xlog.Logger
	Pprof  bool
}
