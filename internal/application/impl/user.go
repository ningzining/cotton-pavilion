package application

import (
	"fmt"
	"user-center/internal/application"
	"user-center/internal/domain/entity"
	"user-center/internal/domain/repository"
	"user-center/internal/utils/crypto"
)

type UserApplication struct {
	UserRepo repository.IUserRepo
}

func (u *UserApplication) Register(dto application.RegisterDTO) error {
	password := crypto.Md5(fmt.Sprintf("%s%s", dto.Mobile, dto.Password))
	user := &entity.User{
		Username: dto.Username,
		Mobile:   dto.Mobile,
		Email:    dto.Email,
		Password: password,
	}
	if err := u.UserRepo.Save(user); err != nil {
		return err
	}

	return nil
}
