package persistence

import (
	"gorm.io/gorm"
	"user-center/internal/domain/entity"
	"user-center/internal/domain/repository"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

func (u UserRepository) Save(user *entity.User) error {
	return u.DB.Save(user).Error
}

func (u UserRepository) FindByMobile(mobile string) (*entity.User, error) {
	var user *entity.User
	err := u.DB.Where("mobile = ?", mobile).Find(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u UserRepository) FindById(userId uint64) (*entity.User, error) {
	var user *entity.User
	err := u.DB.Where("id = ?", userId).Find(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

var _ repository.IUserRepository = &UserRepository{}
