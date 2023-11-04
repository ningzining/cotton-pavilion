package service

import (
	"github.com/google/uuid"
	"strings"
	"time"
	"user-center/internal/domain/entity/do"
	"user-center/internal/enum"
	"user-center/internal/infrastructure/cache"
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
		ExpiredAt: time.Now().Add(time.Minute * 30),
	}
	cache.Save(ticket, value)
	return value
}

func (q QrCodeService) GetQrCode(ticket string) do.QrCode {
	code, ok := cache.Get(ticket).(do.QrCode)
	if !ok {
		code = q.GenerateAndSaveQrCode()
	}
	return code
}

func (q QrCodeService) RefreshQrCode(qrCode do.QrCode) {
	cache.Save(qrCode.Ticket, qrCode)
}

func (q QrCodeService) RemoveQrCode(ticket string) {
	cache.Remove(ticket)
}
