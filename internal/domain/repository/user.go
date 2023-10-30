package repository

import (
	"user-center/internal/domain/entity"
)

type IUserRepository interface {
	Save(user *entity.User) error
	FindByMobile(mobile string) (*entity.User, error)
	FindById(userId uint64) (*entity.User, error)
}
