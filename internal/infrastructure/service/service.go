package service

import "user-center/internal/domain/service"

type Service interface {
	QrCodeService() service.IQrCodeService
}

type svc struct {
}

func NewService() Service {
	return &svc{}
}

func (s *svc) QrCodeService() service.IQrCodeService {
	return service.NewQrCodeService()
}
