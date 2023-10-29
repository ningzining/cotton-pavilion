package repository

import (
	"user-center/internal/domain/entity"
)

type IUserRepo interface {
	Save(user *entity.User) error
	FindByMobile(mobile string) (*entity.User, error)
}
