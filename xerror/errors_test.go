package xerror

import (
	"errors"
	"testing"
)

func TestRepoError_Error(t *testing.T) {
	type fields struct {
		Code    uint32
		Message string
		Err     error
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			"simpleWrapTest",
			fields{100, "message", errors.New("reason")},
			"Error. Code: [100] Message: [message] Reason: [reason]",
		},
		{
			"doubleWrapTest",
			fields{100, "message", &CommonError{
				Code:    200,
				Message: "mid reason",
				Err:     errors.New("reason"),
			}},
			"Error. Code: [100] Message: [message] Reason: [Error. Code: [200] Message: [mid reason] Reason: [reason]]",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := CommonError{
				Code:    tt.fields.Code,
				Message: tt.fields.Message,
				Err:     tt.fields.Err,
			}
			if got := c.Error(); got != tt.want {
				t.Errorf("Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepoError_Unwrap(t *testing.T) {
	type fields struct {
		Code    uint32
		Message string
		Err     error
	}
	errR1 := errors.New("reason")
	errR2 := &CommonError{
		Code:    200,
		Message: "mid reason",
		Err:     errors.New("reason"),
	}
	errR3 := &CommonError{
		Code:    203,
		Message: "mid reason3",
	}
	tests := []struct {
		name    string
		fields  fields
		resErr  error
		wantErr bool
	}{
		{
			"IsTest1",
			fields{100, "message", errR1},
			errR1,
			false,
		},
		{
			"IsTest2",
			fields{100, "message", errR2},
			errR2,
			false,
		},
		{
			"IsTest3",
			fields{100, "message", errR1},
			errR2,
			true,
		},
		{
			"IsTest4",
			fields{203, "mid reason3", errR1},
			errR3,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &CommonError{
				Code:    tt.fields.Code,
				Message: tt.fields.Message,
				Err:     tt.fields.Err,
			}
			if errors.Is(e, tt.resErr) == tt.wantErr {
				t.Errorf("Unwrap() error = %v, is not %v", e, tt.resErr)
			}
		})
	}
}

//func TestTrash(t *testing.T){
//	err1 := errors.New("lol")
//	err2 := errors.New("lol")
//	fmt.Println(err1 == err2)
//
//}

func TestRepoError_Wrap(t *testing.T) {
	errR1 := errors.New("reason1")
	errR2 := &CommonError{
		Code:    202,
		Message: "reason2",
	}
	err1 := &CommonError{
		Code:    500,
		Message: "final error",
	}
	tests := []struct {
		name     string
		fields   *CommonError
		wrapErr  error
		checkErr error
		wantErr  bool
	}{
		{
			"WrapTest1",
			err1,
			errR1,
			errR1,
			true,
		},
		{
			"WrapTest2",
			err1,
			errR2,
			errR2,
			true,
		},
		{
			"WrapTest3",
			err1,
			errR2,
			errR1,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.fields.Wrap(tt.wrapErr); errors.Is(err, tt.checkErr) && errors.Is(err, err1) != tt.wantErr {
				t.Errorf("Wrap() error = %v, is not %v", err, tt.checkErr)
			}
		})
	}
}
