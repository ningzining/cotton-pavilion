package application

import "user-center/internal/infrastructure/persistence"

type Application struct {
	IUserApplication
}

func New(repo *persistence.Repositories) *Application {
	return &Application{
		IUserApplication: &UserApplication{UserRepo: repo.User},
	}
}
