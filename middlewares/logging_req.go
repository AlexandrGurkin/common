package middlewares

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"reflect"
	"regexp"
	"strconv"
	"time"

	"github.com/AlexandrGurkin/common/consts"
	"github.com/AlexandrGurkin/common/xlog"
	guuid "github.com/google/uuid"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	urlReplaceRegexp = regexp.MustCompile(`(v0-9)?(\d+,?)+`)
	summary          = promauto.NewSummaryVec( // nolint:gochecknoglobals
		prometheus.SummaryOpts{
			Name:       "http_response_time",
			Help:       "summary of response time",
			Subsystem:  consts.SubsystemName,
			Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
		}, []string{"route", "method", "status"},
	)
)

const reqLogModule = "log_middleware"
const reqAction = "req_handling"

func RequestLog(next http.Handler, cfg MiddlewareConfig) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log := cfg.Logger
		dump, err := httputil.DumpRequest(r, true)
		if err != nil {
			http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
			return
		}

		cid := r.Header.Get(consts.CorrelationID)
		if len(cid) == 0 {
			cid = guuid.New().String()
			r.Header.Set(consts.CorrelationID, cid)
		}
		logger := log.WithXFields(xlog.Fields{
			consts.FieldModule:        reqLogModule,
			consts.FieldAction:        reqAction,
			consts.FieldCorrelationID: cid,
			consts.FieldURI:           r.RequestURI})

		logger.WithXField(consts.FieldParams, string(dump)).Tracef("Called [%s] method", r.Method)

		startTime := time.Now()

		w.Header().Set(consts.CorrelationID, cid)
		next.ServeHTTP(w, r)
		duration := time.Since(startTime)

		status := int64(0)
		if rw := reflect.Indirect(reflect.ValueOf(w)); rw.IsValid() && rw.Kind() == reflect.Struct {
			if rf := rw.FieldByName("status"); rf.IsValid() && rf.Kind() == reflect.Int {
				status = rf.Int()
			}
		}

		if status != 200 && status != 204 {
			logger.WithXFields(xlog.Fields{
				consts.FieldHttpCode: status,
				consts.FieldDuration: duration.Microseconds(),
				consts.FieldParams:   string(dump),
			}).Warnf("Completed [%s] method", r.Method)
		} else {
			logger.WithXFields(xlog.Fields{
				consts.FieldHttpCode: status,
				consts.FieldDuration: duration.Microseconds(),
			}).Tracef("Completed [%s] method", r.Method)
		}

		summary.WithLabelValues(asteriskRequestRoute(r.RequestURI), r.Method,
			fmt.Sprintf("%sxx", strconv.FormatInt(status, 10)[:1])).
			Observe(duration.Seconds() * 1000)
	})
}

func asteriskRequestRoute(requestURI string) string {
	return urlReplaceRegexp.ReplaceAllString(requestURI, "*")
}
