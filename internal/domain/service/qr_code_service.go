package service

import (
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"strings"
	"time"
	"user-center/internal/domain/entity/do"
	"user-center/internal/enum"
	"user-center/internal/infrastructure/cache/qr_code_cache"
	"user-center/internal/infrastructure/cache/qr_code_conn_cache"
)

type QrCodeService struct{}

func NewQrCodeService() IQrCodeService {
	return &QrCodeService{}
}

type IQrCodeService interface {
	GenerateAndSaveQrCode() do.QrCode
	GetQrCode(ticket string) do.QrCode
	SaveQrCode(qrCode do.QrCode)
	RemoveTicket(ticket string)

	GetTicket(conn *websocket.Conn) string
	SaveConn(conn *websocket.Conn, ticket string)
	RemoveConn(conn *websocket.Conn)
}

func (q QrCodeService) GenerateAndSaveQrCode() do.QrCode {
	ticket := strings.ReplaceAll(uuid.New().String(), "-", "")
	value := do.QrCode{
		Ticket:    ticket,
		Status:    enum.QrCodeStatusUnauthorized,
		ExpiredAt: time.Now().Add(time.Second * 30),
	}
	qr_code_cache.Save(ticket, value)
	return value
}

func (q QrCodeService) GetQrCode(ticket string) do.QrCode {
	qrCode, ok := qr_code_cache.Get(ticket).(do.QrCode)
	if !ok {
		qrCode = q.GenerateAndSaveQrCode()
	}
	if qrCode.IsExpired() {
		q.RemoveTicket(qrCode.Ticket)
		qrCode = q.GenerateAndSaveQrCode()
	}
	return qrCode
}

func (q QrCodeService) SaveQrCode(qrCode do.QrCode) {
	qr_code_cache.Save(qrCode.Ticket, qrCode)
}

func (q QrCodeService) RemoveTicket(ticket string) {
	qr_code_cache.Remove(ticket)
}

func (q QrCodeService) GetTicket(conn *websocket.Conn) string {
	return qr_code_conn_cache.Get(conn)
}

func (q QrCodeService) SaveConn(conn *websocket.Conn, ticket string) {
	qr_code_conn_cache.Save(conn, ticket)
}

func (q QrCodeService) RemoveConn(conn *websocket.Conn) {
	qr_code_conn_cache.Remove(conn)
}
