package mysql

import (
	"github.com/ningzining/cotton-pavilion/internal/domain/model"
	"github.com/ningzining/cotton-pavilion/internal/domain/repository"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(repository Repository) repository.UserRepository {
	return &UserRepository{
		db: repository.DB,
	}
}

func (u UserRepository) Save(user *model.User) error {
	return u.db.Save(&user).Error
}

func (u UserRepository) FindByMobile(mobile string) (*model.User, error) {
	var user *model.User
	err := u.db.Where("mobile = ?", mobile).Find(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u UserRepository) FindById(userId uint64) (*model.User, error) {
	var user *model.User
	err := u.db.Where("id = ?", userId).Find(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}
