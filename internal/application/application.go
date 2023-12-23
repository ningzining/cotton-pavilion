package application

import (
	"github.com/ningzining/cotton-pavilion/internal/domain/service"
	"github.com/ningzining/cotton-pavilion/internal/infrastructure/store"
)

type Application interface {
	UserApplication() UserApplication
	ImageApplication() ImageApplication
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

func (a app) UserApplication() UserApplication {
	return newUserApplication(a.Store, a.Service)
}

func (a app) ImageApplication() ImageApplication {
	return newImageApplication(a.Store)
}
