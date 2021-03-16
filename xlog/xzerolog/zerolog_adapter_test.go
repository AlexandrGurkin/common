package xzerolog

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/AlexandrGurkin/common/mocks"
	"github.com/AlexandrGurkin/common/xlog"
	"github.com/golang/mock/gomock"
)

func BenchmarkXLog(b *testing.B) {
	log := NewXZerolog(xlog.LoggerCfg{Level: "trace", Out: &xlog.BlackholeStream{}})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		log.Trace("ep")
	}
}

func TestNewXZerologSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	expectMsg := []byte("{\"level\":\"trace\",\"message\":\"ep\",\"time\":\"*\"}")

	mockWriter := mocks.NewMockWriter(ctrl)
	mockWriter.EXPECT().
		Write(EqWriter(expectMsg)).
		Times(1)

	log := NewXZerolog(xlog.LoggerCfg{Level: "trace", Out: mockWriter})
	log.Trace("ep")
} //{"level":"trace","msg":"ep","time":"2021-02-27T16:07:06.536258+03:00"}

func Test_x_WithXField(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	expectMsg1 := []byte("{\"level\":\"trace\",\"message\":\"ep\",\"time\":\"*\"}")
	expectMsg2 := []byte("{\"key\":\"value\",\"level\":\"trace\",\"message\":\"ep\",\"time\":\"*\"}")
	mockWriter := mocks.NewMockWriter(ctrl)

	log1 := NewXZerolog(xlog.LoggerCfg{Level: "trace", Out: mockWriter})
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

	expectMsg1 := []byte("{\"level\":\"trace\",\"message\":\"ep\",\"time\":\"*\"}")
	expectMsg2 := []byte("{\"key\":\"value\",\"level\":\"trace\",\"lol\":\"kek\",\"message\":\"ep\",\"time\":\"*\"}")
	mockWriter := mocks.NewMockWriter(ctrl)

	log1 := NewXZerolog(xlog.LoggerCfg{Level: "trace", Out: mockWriter})
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
