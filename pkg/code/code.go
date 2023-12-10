package code

import (
	"user-center/pkg/errors"
)

type ErrCode struct {
	C    int
	HTTP int
	Ext  string
}

func (e *ErrCode) Code() int {
	return e.C
}

func (e *ErrCode) HTTPStatus() int {
	return e.HTTP
}

func (e *ErrCode) String() string {
	return e.Ext
}

func register(code, httpStatus int, message string) {
	coder := &ErrCode{
		C:    code,
		HTTP: httpStatus,
		Ext:  message,
	}
	errors.MustRegister(coder)
}
