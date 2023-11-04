package persistence

import (
	"gorm.io/gorm"
	"user-center/internal/domain/entity"
	"user-center/internal/domain/repository"
)

type UserRepository struct {
	DB *gorm.DB
}

func (u UserRepository) Save(user *entity.User) error {
	if user.ID > 0 {
		return u.DB.Updates(user).Error
	}
	return u.DB.Create(user).Error
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

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

var _ repository.IUserRepository = &UserRepository{}