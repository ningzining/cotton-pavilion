package application

import (
	"user-center/internal/domain/service"
	"user-center/internal/infrastructure/cache"
	"user-center/internal/infrastructure/persistence"
)

type Application struct {
	Cache   cache.ICache
	UserApp IUserApplication
}

func New(repo *persistence.Repositories, svc *service.Services, cache cache.ICache) *Application {
	return &Application{
		Cache:   cache,
		UserApp: NewUserApplication(repo.User, svc.QrCodeService),
	}
}
