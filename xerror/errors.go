// Package xerror contains custom implementations Error interface.
package xerror

import (
	"fmt"
)

// CommonError is custom error type.
type CommonError struct {
	Code    uint32
	Message string
	Err     error
}

// Error is implementation for custom type.
func (e *CommonError) Error() string {
	reason := "nil"
	if e.Err != nil {
		reason = e.Err.Error()
	}

	return fmt.Sprintf("Error. Code: [%d] Message: [%s] Reason: [%s]", e.Code, e.Message, reason)
}

// Unwrap is implementation for custom type.
func (e *CommonError) Unwrap() error { return e.Err }

// Wrap is implementation for custom type.
func (e *CommonError) Wrap(er error) error {
	return &CommonError{
		Code:    e.Code,
		Message: e.Message,
		Err:     er,
	}
}

// Wrapf is implementation for custom type.
//nolint: goerr113
func (e *CommonError) Wrapf(format string, args ...interface{}) error {
	return &CommonError{
		Code:    e.Code,
		Message: e.Message,
		Err:     fmt.Errorf(format, args...),
	}
}

// Is is implementation for custom type.
func (e *CommonError) Is(target error) bool {
	t, ok := target.(*CommonError)
	if !ok {
		return false
	}

	return (e.Code == t.Code) && (e.Message == t.Message)
}
