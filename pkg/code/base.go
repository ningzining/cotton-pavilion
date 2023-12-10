package code

import "net/http"

// 错误定义:
// 10: 服务
// 00: 模块
// 01: 错误描述
// 通用异常
const (
	ErrSuccess = iota + 100001
	ErrUnKnow
	ErrBind
	ErrValidation
	ErrTokenInvalid
	ErrPageNotFind
)

// ErrDatabase 数据库操作异常
const (
	ErrDatabase int = iota + 100101
)

// 授权异常
const (
	ErrPasswordIncorrect int = iota + 100201
	ErrQrCodeExpired
	ErrTokenGenerate
	ErrTokenExpired
	ErrMissingHeader
	ErrMissingToken
	ErrPermissionDenied
)

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

	// 通用模块错误信息
	register(ErrPasswordIncorrect, http.StatusOK, "密码不正确,请重新输入")
	register(ErrQrCodeExpired, http.StatusUnauthorized, "二维码已过期,请刷新")
	register(ErrTokenGenerate, http.StatusOK, "token生成失败")
	register(ErrMissingHeader, http.StatusUnauthorized, "缺少header信息，请重新登录")
	register(ErrMissingToken, http.StatusUnauthorized, "header中缺少token信息，请重新登录")
	register(ErrTokenExpired, http.StatusUnauthorized, "token信息已过期，请重新登录")
	register(ErrPermissionDenied, http.StatusForbidden, "权限不存在，请联系管理员")
}
