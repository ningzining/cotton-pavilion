package service

import (
	"user-center/internal/domain/service"
)

type Service struct {
	QrCodeService service.IQrCodeService
}

func New() *Service {
	return &Service{
		QrCodeService: service.NewQrCodeService(),
	}
}
