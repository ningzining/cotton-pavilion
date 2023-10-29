package persistence

import (
	"gorm.io/gorm"
	"user-center/internal/domain/entity"
	"user-center/internal/domain/repository"
)

type UserRepo struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepo {
	return &UserRepo{
		DB: db,
	}
}

var _ repository.IUserRepo = &UserRepo{}

func (u *UserRepo) Save(user *entity.User) error {
	if user.ID > 0 {
		return u.DB.Updates(user).Error
	}
	return u.DB.Create(user).Error
}

func (u *UserRepo) FindByMobile(mobile string) (*entity.User, error) {
	var user *entity.User
	err := u.DB.Where("mobile = ?", mobile).Find(&user).Error
	return user, err
}
