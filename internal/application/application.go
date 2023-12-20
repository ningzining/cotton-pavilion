package application

import (
	"user-center/internal/domain/service"
	"user-center/internal/infrastructure/store"
	"user-center/internal/infrastructure/third_party"
)

type Application interface {
	UserApplication() UserApplication
	ImageApplication() ImageApplication
}

type app struct {
	Store   store.Factory
	Service service.Service
	Oss     third_party.Oss
}

func NewApplication(store store.Factory, service service.Service, oss third_party.Oss) Application {
	return &app{
		Store:   store,
		Service: service,
		Oss:     oss,
	}
}

func (a app) UserApplication() UserApplication {
	return newUserApplication(a.Store, a.Service)
}

func (a app) ImageApplication() ImageApplication {
	return newImageApplication(a.Store, a.Oss)
}
