package impl

import (
	"user-center/internal/domain/repository"
	"user-center/internal/infrastructure/db/mysql"
)

type Repository struct {
	repository.IUserRepo
}

func New() *Repository {
	return &Repository{
		IUserRepo: &UserRepo{DB: mysql.DB()},
	}
}
