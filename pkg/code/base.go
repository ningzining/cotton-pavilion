package code

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
	ErrPermissionDenied
)
