package service

import "user-center/internal/infrastructure/store"

type Service interface {
	QrCodeService() QrCodeService
}

type svc struct {
	store store.Factory
}

func NewService(store store.Factory) Service {
	return &svc{
		store: store,
	}
}

func (s *svc) QrCodeService() QrCodeService {
	return newQrCodeService()
}
