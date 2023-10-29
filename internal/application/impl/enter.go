package application

import (
	"user-center/internal/application"
	"user-center/internal/infrastructure/persistence"
)

type Application struct {
	application.IUserApplication
}

func New(repo *persistence.Repositories) *Application {
	return &Application{
		IUserApplication: &UserApplication{UserRepo: repo.User},
	}
}
