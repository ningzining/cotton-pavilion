package application

import (
	"user-center/internal/domain/service"
	"user-center/internal/infrastructure/store"
)

type Application interface {
	UserApplication() UserApplication
}

type app struct {
	Store   store.Factory
	Service service.Service
}

func NewApplication(store store.Factory, service service.Service) Application {
	return &app{
		Store:   store,
		Service: service,
	}
}

func (a *app) UserApplication() UserApplication {
	return newUserApplication(a.Store, a.Service)
}
