package application

import "github.com/gorilla/websocket"

type RegisterDTO struct {
	Username string `json:"username" binding:"required"`
	Mobile   string `json:"mobile" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginDTO struct {
	Mobile   string `json:"mobile" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginRet struct {
	Token string `json:"token"`
}

type QrCodeStatus = string

const (
	QrCodeStatusUnauthorized QrCodeStatus = "UNAUTHORIZED" // 未授权
	QrCodeStatusAuthorized   QrCodeStatus = "AUTHORIZED"   // 已授权
)

type QrCodeDTO struct {
	Conn   *websocket.Conn
	Ticket string
}

type QrCodeRet struct {
	Ticket string       `json:"ticket"`
	Status QrCodeStatus `json:"status"`
	Token  string       `json:"token"`
}

type ConfirmLoginDTO struct {
	TemporaryToken string `json:"temporary_token"` // todo 待实现二维码已扫描的功能
	Ticket         string `json:"ticket"`
}

type IUserApplication interface {
	Register(dto RegisterDTO) error
	Login(dto LoginDTO) (*LoginRet, error)
	QrCode(dto QrCodeDTO) (*QrCodeRet, error)
	ConfirmLogin(dto ConfirmLoginDTO) error
}
