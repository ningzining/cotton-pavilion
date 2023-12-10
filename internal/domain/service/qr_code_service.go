package service

import (
	"errors"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"strings"
	"time"
	"user-center/internal/domain/entity/do"
	"user-center/internal/domain/entity/enum"
	"user-center/internal/infrastructure/cache/qr_code_conn_cache"
	"user-center/internal/infrastructure/cache/qr_code_info_cache"
)

type QrCodeService struct{}

func NewQrCodeService() IQrCodeService {
	return &QrCodeService{}
}

type IQrCodeService interface {
	GenerateNew() *do.QrCode
	GetTicket(conn *websocket.Conn) string
	GetQrCode(ticket string) (*do.QrCode, error)
	Remove(conn *websocket.Conn, ticket string)
	SaveQrCode(qrCode *do.QrCode)
}

func (q QrCodeService) GetTicket(conn *websocket.Conn) string {
	var resTicket string
	ticket, b := qr_code_conn_cache.Get(conn)
	if !b {
		newQrCode := q.GenerateNew()
		// 先保存连接和ticket的关系，再保存ticket和具体的二维码信息
		qr_code_conn_cache.Save(conn, newQrCode.Ticket)
		qr_code_info_cache.Save(newQrCode.Ticket, newQrCode)
		resTicket = newQrCode.Ticket
	} else {
		resTicket = ticket.(string)
	}
	return resTicket
}

func (q QrCodeService) GenerateNew() *do.QrCode {
	ticket := strings.ReplaceAll(uuid.New().String(), "-", "")
	return &do.QrCode{
		Ticket:    ticket,
		Status:    enum.QrCodeStatusUnauthorized,
		ExpiredAt: time.Now().Add(time.Second * 30),
	}
}

func (q QrCodeService) GetQrCode(ticket string) (*do.QrCode, error) {
	qrCode, ok := qr_code_info_cache.Get(ticket)
	if !ok {
		return nil, errors.New("二维码不存在")
	}
	code, ok := qrCode.(*do.QrCode)
	if !ok {
		return nil, errors.New("二维码异常")
	}
	return code, nil
}

// Remove 删除二维码信息和连接的信息
func (q QrCodeService) Remove(conn *websocket.Conn, ticket string) {
	qr_code_conn_cache.Remove(conn)
	qr_code_info_cache.Remove(ticket)
}

func (q QrCodeService) SaveQrCode(qrCode *do.QrCode) {
	qr_code_info_cache.Save(qrCode.Ticket, qrCode)
}
