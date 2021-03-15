package middlewares

import (
	"net/http"
	"net/http/pprof"
)

const (
	pprofUrl        string = `/debug/pprof/`
	pprofCmdUrl     string = `/debug/pprof/cmdline`
	pprofProfileUrl string = `/debug/pprof/profile`
	pprofSymbolUrl  string = `/debug/pprof/symbol`
	pprofTraceUrl   string = `/debug/pprof/trace`
	pprofAllocsUrl  string = `/debug/pprof/allocs`
	pprofHeapUrl    string = `/debug/pprof/heap`
)

func PprofHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		switch req.URL.Path {
		case pprofUrl:
			pprof.Index(w, req)
			return
		case pprofCmdUrl:
			pprof.Cmdline(w, req)
			return
		case pprofProfileUrl:
			pprof.Profile(w, req)
			return
		case pprofSymbolUrl:
			pprof.Profile(w, req)
			return
		case pprofTraceUrl:
			pprof.Trace(w, req)
			return
		case pprofAllocsUrl:
			pprof.Handler("allocs").ServeHTTP(w, req)
			return
		case pprofHeapUrl:
			pprof.Handler("heap").ServeHTTP(w, req)
			return
		default:
			next.ServeHTTP(w, req)
		}
	})
}
