package application

import (
	"user-center/internal/application"
	"user-center/internal/domain/repository/impl"
)

type Application struct {
	application.IUserApplication
}

func New(repo *impl.Repository) *Application {
	return &Application{
		IUserApplication: &UserApplication{UserRepo: repo.IUserRepo},
	}
}
