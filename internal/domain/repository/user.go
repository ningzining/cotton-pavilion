package repository

import (
	"user-center/internal/domain/model"
)

type UserRepository interface {
	Save(user *model.User) error
	FindByMobile(mobile string) (*model.User, error)
	FindById(userId uint64) (*model.User, error)
}
