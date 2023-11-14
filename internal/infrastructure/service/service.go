package service

import (
	"user-center/internal/domain/service"
	"user-center/internal/infrastructure/persistence"
)

type Service struct {
	QrCodeService service.IQrCodeService
}

func New(repositories *persistence.Repositories) *Service {
	return &Service{
		QrCodeService: service.NewQrCodeService(),
	}
}
