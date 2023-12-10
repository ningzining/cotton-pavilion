package types

import (
	"github.com/gorilla/websocket"
	"user-center/internal/domain/entity/enum"
)

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

type QrCodeDTO struct {
	Conn *websocket.Conn
}

type QrCodeRet struct {
	Ticket string            `json:"ticket"`
	Status enum.QrCodeStatus `json:"status"`
	Token  string            `json:"token"`
}

type ConfirmLoginDTO struct {
	TemporaryToken string `form:"temporary_token" binding:"required"` // 临时token
	Ticket         string `form:"ticket" binding:"required"`          // 二维码
	Token          string `header:"Authorization"`                    // 用户token
}

type ScanQrCodeDTO struct {
	Ticket string `form:"ticket" binding:"required"` // 二维码
	Token  string `header:"Authorization"`           // token
}

type ScanQrCodeRet struct {
	TemporaryToken string `json:"temporary_token"`
}
