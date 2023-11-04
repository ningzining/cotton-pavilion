package service

import (
	"github.com/google/uuid"
	"strings"
	"time"
	"user-center/internal/domain/entity/do"
	"user-center/internal/enum"
	"user-center/internal/infrastructure/cache/qr_code_cache"
)

type QrCodeService struct{}

func NewQrCodeService() IQrCodeService {
	return &QrCodeService{}
}

type IQrCodeService interface {
	GenerateAndSaveQrCode() do.QrCode
	GetQrCode(ticket string) do.QrCode
	RefreshQrCode(qrCode do.QrCode)
	RemoveQrCode(ticket string)
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
	code, ok := qr_code_cache.Get(ticket).(do.QrCode)
	if !ok {
		code = q.GenerateAndSaveQrCode()
	}
	return code
}

func (q QrCodeService) RefreshQrCode(qrCode do.QrCode) {
	qr_code_cache.Save(qrCode.Ticket, qrCode)
}

func (q QrCodeService) RemoveQrCode(ticket string) {
	qr_code_cache.Remove(ticket)
}
