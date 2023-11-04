package service

type Services struct {
	QrCodeService IQrCodeService
}

func New() *Services {
	return &Services{
		QrCodeService: NewQrCodeService(),
	}
}
