package application

import (
	"errors"
	"fmt"
	"user-center/internal/domain/entity"
	"user-center/internal/domain/repository"
	"user-center/internal/utils/crypto"
)

type UserApplication struct {
	UserRepo repository.IUserRepository
}

func (u *UserApplication) Login(dto LoginDTO) (*LoginRet, error) {
	password := crypto.Md5(fmt.Sprintf("%s%s", dto.Mobile, dto.Password))
	user, err := u.UserRepo.FindByMobile(dto.Mobile)
	if err != nil {
		return nil, err
	}
	if user.Password != password {
		return nil, errors.New("登录的账号或密码不正确")
	}
	return &LoginRet{
		Token: "testToken",
	}, nil
}

func (u *UserApplication) Register(dto RegisterDTO) error {
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

var _ IUserApplication = &UserApplication{}
