package application

import (
	"user-center/internal/infrastructure/service"
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
	return NewUserApplication(a.Store, a.Service)
}
