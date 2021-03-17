package mocks

import "sync/atomic"

type BlackHoleStream struct {
	writeCount uint64
}

func (s *BlackHoleStream) WriteCount() uint64 {
	return atomic.LoadUint64(&s.writeCount)
}

func (s *BlackHoleStream) Write(p []byte) (int, error) {
	atomic.AddUint64(&s.writeCount, 1)
	return len(p), nil
}
