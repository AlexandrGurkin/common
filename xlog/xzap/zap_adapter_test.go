package xzap

import (
	"testing"
	"time"

	"github.com/AlexandrGurkin/common/mocks"
	"github.com/AlexandrGurkin/common/xlog"
	"github.com/golang/mock/gomock"
)

func BenchmarkXLog(b *testing.B) {
	log, _ := NewXZap(xlog.LoggerCfg{Level: "trace", Out: &mocks.BlackHoleStream{}, TimeFormat: time.RFC3339Nano})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		log.Trace("ep")
	}
}

func BenchmarkXLogWithField(b *testing.B) {
	log, _ := NewXZap(xlog.LoggerCfg{Level: "trace", Out: &mocks.BlackHoleStream{}, TimeFormat: time.RFC3339Nano})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		log := log.WithXField("key", "value")
		log.Trace("ep")
	}
}

func BenchmarkXLogWithFields(b *testing.B) {
	log, _ := NewXZap(xlog.LoggerCfg{Level: "trace", Out: &mocks.BlackHoleStream{}, TimeFormat: time.RFC3339Nano})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		log := log.WithXFields(xlog.Fields{"key1": "val1", "key2": "val2"})
		log.Trace("ep")
	}
}

func BenchmarkXLog_f(b *testing.B) {
	log, _ := NewXZap(xlog.LoggerCfg{Level: "trace", Out: &mocks.BlackHoleStream{}, TimeFormat: time.RFC3339Nano})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		log.Tracef("ep %s", "1")
	}
}

func BenchmarkXLogWithField_f(b *testing.B) {
	log, _ := NewXZap(xlog.LoggerCfg{Level: "trace", Out: &mocks.BlackHoleStream{}, TimeFormat: time.RFC3339Nano})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		log := log.WithXField("key", "value")
		log.Tracef("ep %s", "1")
	}
}

func BenchmarkXLogWithFields_f(b *testing.B) {
	log, _ := NewXZap(xlog.LoggerCfg{Level: "trace", Out: &mocks.BlackHoleStream{}, TimeFormat: time.RFC3339Nano})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		log := log.WithXFields(xlog.Fields{"key1": "val1", "key2": "val2"})
		log.Tracef("ep %s", "1")
	}
}

//For zap trace level eq debug level
func TestNewXZapSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	expectMsg := []byte("{\"level\":\"debug\",\"message\":\"ep\",\"time\":\"*\"}")

	mockWriter := mocks.NewMockWriter(ctrl)
	mockWriter.EXPECT().
		Write(mocks.EqWriter(expectMsg)).
		Times(1)

	log, err := NewXZap(xlog.LoggerCfg{Level: "trace", Out: mockWriter, TimeFormat: time.RFC3339Nano})
	if err != nil {
		t.Error(err)
	}
	log.Trace("ep")
} //{"level":"trace","msg":"ep","time":"2021-02-27T16:07:06.536258+03:00"}

func Test_x_WithXField(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	expectMsg1 := []byte("{\"level\":\"debug\",\"message\":\"ep\",\"time\":\"*\"}")
	expectMsg2 := []byte("{\"key\":\"value\",\"level\":\"debug\",\"message\":\"ep\",\"time\":\"*\"}")
	mockWriter := mocks.NewMockWriter(ctrl)

	log1, err := NewXZap(xlog.LoggerCfg{Level: "trace", Out: mockWriter, TimeFormat: time.RFC3339Nano})
	if err != nil {
		t.Error(err)
	}
	mockWriter.EXPECT().
		Write(mocks.EqWriter(expectMsg1)).
		Times(1)
	log1.Trace("ep")

	log2 := log1.WithXField("key", "value")
	mockWriter.EXPECT().
		Write(mocks.EqWriter(expectMsg2)).
		Times(1)
	log2.Tracef("ep")

	//like first
	mockWriter.EXPECT().
		Write(mocks.EqWriter(expectMsg1)).
		Times(1)
	log1.Trace("ep")

} //{"key":"value","level":"trace","msg":"ep","time":"2021-02-28T18:09:27.889963+03:00"}

func Test_x_WithXFields(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	expectMsg1 := []byte("{\"level\":\"debug\",\"message\":\"ep\",\"time\":\"*\"}")
	expectMsg2 := []byte("{\"key\":\"value\",\"level\":\"debug\",\"lol\":\"kek\",\"message\":\"ep\",\"time\":\"*\"}")
	mockWriter := mocks.NewMockWriter(ctrl)

	log1, _ := NewXZap(xlog.LoggerCfg{Level: "trace", Out: mockWriter, TimeFormat: time.RFC3339Nano})
	mockWriter.EXPECT().
		Write(mocks.EqWriter(expectMsg1)).
		Times(1)
	log1.Trace("ep")

	log2 := log1.WithXFields(xlog.Fields{"key": "value", "lol": "kek"})
	mockWriter.EXPECT().
		Write(mocks.EqWriter(expectMsg2)).
		Times(1)
	log2.Trace("ep")

	//like first
	mockWriter.EXPECT().
		Write(mocks.EqWriter(expectMsg1)).
		Times(1)
	log1.Trace("ep")
} //{"key":"value","level":"trace","lol":"kek","msg":"ep","time":"2021-02-28T18:48:33.626854+03:00"}
