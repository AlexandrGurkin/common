package mocks

import (
	"encoding/json"
	"fmt"
	"sync/atomic"

	"github.com/golang/mock/gomock"
)

type BlackHoleStream struct {
	writeCount uint64
	syncCount  uint64
}

func (s *BlackHoleStream) WriteCount() uint64 {
	return atomic.LoadUint64(&s.writeCount)
}

func (s *BlackHoleStream) SyncCount() uint64 {
	return atomic.LoadUint64(&s.syncCount)
}

func (s *BlackHoleStream) Write(p []byte) (int, error) {
	atomic.AddUint64(&s.writeCount, 1)
	return len(p), nil
}

func (s *BlackHoleStream) Sync() error {
	atomic.AddUint64(&s.syncCount, 1)
	return nil
}

func EqWriter(t []byte) gomock.Matcher {
	var f map[string]interface{}
	_ = json.Unmarshal(t, &f)
	return &eqWriter{f}
}

type eqWriter struct{ t map[string]interface{} }

func (o *eqWriter) Matches(x interface{}) bool {
	if val, ok := x.([]byte); ok {
		var r map[string]interface{}
		_ = json.Unmarshal(val, &r)
		for k, v := range o.t {
			if k == "time" {
				continue
			}
			if v != r[k] {
				return false
			}
		}
		return true
	}
	return false
}

func (o *eqWriter) String() string {
	return fmt.Sprintf("equal [%v]", o.t)
}
