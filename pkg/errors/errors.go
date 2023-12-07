package errors

import "fmt"

type withCode struct {
	err  error
	code int
}

func (w *withCode) Error() string { return fmt.Sprintf("%s", w.err.Error()) }

func WithCode(code int, format string, args ...interface{}) error {
	return &withCode{
		err:  fmt.Errorf(format, args...),
		code: code,
	}
}
