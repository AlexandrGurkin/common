package xlogrus

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/AlexandrGurkin/common/mocks"
	"github.com/AlexandrGurkin/common/xlog"
	"github.com/golang/mock/gomock"
)

func BenchmarkXLog(b *testing.B) {
	log := NewXLogrus(xlog.LoggerCfg{Level: "trace", Out: &mocks.BlackHoleStream{}})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		log.Trace("ep")
	}
}

func TestNewXLogrusSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	expectMsg := []byte("{\"level\":\"trace\",\"message\":\"ep\",\"time\":\"")

	mockWriter := mocks.NewMockWriter(ctrl)
	mockWriter.EXPECT().
		Write(EqWriter(expectMsg)).
		Times(1)

	log := NewXLogrus(xlog.LoggerCfg{Level: "trace", Out: mockWriter})
	log.Trace("ep")
} //{"level":"trace","msg":"ep","time":"2021-02-27T16:07:06.536258+03:00"}

func Test_xrus_WithXField(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	expectMsg1 := []byte("{\"level\":\"trace\",\"message\":\"ep\",\"time\":\"")
	expectMsg2 := []byte("{\"key\":\"value\",\"level\":\"trace\",\"message\":\"ep\",\"time\":\"")
	mockWriter := mocks.NewMockWriter(ctrl)

	log1 := NewXLogrus(xlog.LoggerCfg{Level: "trace", Out: mockWriter})
	mockWriter.EXPECT().
		Write(EqWriter(expectMsg1)).
		Times(1)
	log1.Trace("ep")

	log2 := log1.WithXField("key", "value")
	mockWriter.EXPECT().
		Write(EqWriter(expectMsg2)).
		Times(1)
	log2.Trace("ep")

	//like first
	mockWriter.EXPECT().
		Write(EqWriter(expectMsg1)).
		Times(1)
	log1.Trace("ep")

} //{"key":"value","level":"trace","msg":"ep","time":"2021-02-28T18:09:27.889963+03:00"}

func Test_xrus_WithXFields(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	expectMsg1 := []byte("{\"level\":\"trace\",\"message\":\"ep\",\"time\":\"")
	expectMsg2 := []byte("{\"key\":\"value\",\"level\":\"trace\",\"lol\":\"kek\",\"message\":\"ep\",\"time\":\"")
	mockWriter := mocks.NewMockWriter(ctrl)

	log1 := NewXLogrus(xlog.LoggerCfg{Level: "trace", Out: mockWriter})
	mockWriter.EXPECT().
		Write(EqWriter(expectMsg1)).
		Times(1)
	log1.Trace("ep")

	log2 := log1.WithXFields(xlog.Fields{"key": "value", "lol": "kek"})
	mockWriter.EXPECT().
		Write(EqWriter(expectMsg2)).
		Times(1)
	log2.Trace("ep")

	//like first
	mockWriter.EXPECT().
		Write(EqWriter(expectMsg1)).
		Times(1)
	log1.Trace("ep")
} //{"key":"value","level":"trace","lol":"kek","msg":"ep","time":"2021-02-28T18:48:33.626854+03:00"}

func EqWriter(t []byte) gomock.Matcher {
	return &eqWriter{t}
}

type eqWriter struct{ t []byte }

func (o *eqWriter) Matches(x interface{}) bool {
	if val, ok := x.([]byte); ok {
		if len(val) < len(o.t) {
			return false
		}
		tmp := val[:len(o.t)]
		if bytes.Compare(tmp, o.t) == 0 {
			return true
		}
	}
	return false
}

func (o *eqWriter) String() string {
	return fmt.Sprintf("starts with [%v]", o.t)
}
