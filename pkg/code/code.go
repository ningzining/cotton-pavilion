package code

import (
	"net/http"
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

func init() {
	// 基础错误信息
	register(ErrSuccess, http.StatusOK, "OK")
	register(ErrUnKnow, http.StatusOK, "未知异常,请联系管理员")
	register(ErrBind, http.StatusOK, "参数绑定异常")
	register(ErrValidation, http.StatusOK, "参数校验失败")
	register(ErrTokenInvalid, http.StatusUnauthorized, "令牌无效,请重新登录")
	register(ErrPageNotFind, http.StatusNotFound, "接口不存在")

	// 通用错误信息
	register(ErrDatabase, http.StatusOK, "数据库异常")

	// 用户模块错误信息
	register(ErrPasswordIncorrect, http.StatusOK, "密码不正确,请重新输入")
	register(ErrQrCodeExpired, http.StatusUnauthorized, "二维码已过期,请刷新")
	register(ErrTokenGenerate, http.StatusOK, "token生成失败")
}
