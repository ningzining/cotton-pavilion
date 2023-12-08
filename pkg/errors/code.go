package errors

import (
	"fmt"
	"net/http"
	"sync"
)

var (
	unknownCoder = defaultCoder{1, http.StatusInternalServerError, "服务未知异常"}
)

type defaultCoder struct {
	C    int
	HTTP int
	Ext  string
}

func (coder defaultCoder) Code() int {
	return coder.C
}

func (coder defaultCoder) String() string {
	return coder.Ext
}

func (coder defaultCoder) HTTPStatus() int {
	return coder.HTTP
}

type Coder interface {
	Code() int
	HTTPStatus() int
	String() string
}

var codes = map[int]Coder{}
var codeMux = &sync.Mutex{}

func MustRegister(coder Coder) {
	codeMux.Lock()
	defer codeMux.Unlock()
	if _, ok := codes[coder.Code()]; ok {
		panic(fmt.Sprintf("code: %d already exist", coder.Code()))
	}
	codes[coder.Code()] = coder
}

func ParseCoder(err error) Coder {
	if err == nil {
		return nil
	}
	if v, ok := err.(*withCode); ok {
		if coder, ok := codes[v.code]; ok {
			return coder
		}
	}
	return unknownCoder
}

func init() {
	codes[unknownCoder.Code()] = unknownCoder
}
