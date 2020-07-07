package api

import (
	"fmt"
)

// Errorer is implemented by all concrete response types that may contain
// errors. It allows us to change the HTTP response code without needing to
// trigger an endpoint (transport-level) error. For more information, read the
// big comment in endpoints.go.
type Errorer interface {
	GetError() *Error
}

type Error struct {
	Code    int32  `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

// Error implements the error interface.
func (e *Error) Error() string {
	return e.Message
}

// WrapErr returns a corev1.Error for the given error and msg.
func WrapErr(err error, msg string) error {
	if err == nil {
		return nil
	}
	e := &Error{Message: fmt.Sprintf("%s: %s", msg, err.Error())}
	if v1err, ok := err.(*Error); ok {
		e.Code = v1err.Code
	}
	return e
}
