package persistence

import (
	"gorm.io/gorm"
	"user-center/internal/domain/repository"
)

type Repositories struct {
	db   *gorm.DB
	User repository.IUserRepo
}

func NewRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		db:   db,
		User: NewUserRepository(db),
	}
}
