package application

import (
	"user-center/internal/domain/service"
	"user-center/internal/infrastructure/persistence"
)

type Application struct {
	IUserApplication
}

func New(repo *persistence.Repositories, svc *service.Services) *Application {
	return &Application{
		IUserApplication: NewUserApplication(repo.User, svc.QrCodeService),
	}
}
