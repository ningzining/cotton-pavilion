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
	QrCodeStatusUnauthorized QrCodeStatus = "Unauthorized" // 未授权
	QrCodeStatusAuthorized   QrCodeStatus = "Authorized"   // 已授权
)

type QrCodeDTO struct {
	Conn   *websocket.Conn
	Ticket string
}

type QrCodeRet struct {
	Ticket string       `json:"ticket"`
	Status QrCodeStatus `json:"status"`
}

type IUserApplication interface {
	Register(dto RegisterDTO) error
	Login(dto LoginDTO) (*LoginRet, error)
	QrCode(dto QrCodeDTO) *QrCodeRet
}
