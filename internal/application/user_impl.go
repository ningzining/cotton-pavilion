package application

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
	"user-center/internal/consts"
	"user-center/internal/domain/entity"
	"user-center/internal/domain/repository"
	"user-center/internal/infrastructure/cache"
	"user-center/internal/utils/crypto"
	"user-center/internal/utils/jwtutils"
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

	claims := jwtutils.Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:   consts.SystemName,
			Subject:  consts.JwtSubject,
			IssuedAt: jwt.NewNumericDate(time.Now()),
		},
		User: jwtutils.User{
			Id:       user.ID,
			Username: user.Username,
		},
	}
	token, err := jwtutils.GenerateJwt(claims, consts.JwtSecret)
	if err != nil {
		return nil, err
	}

	return &LoginRet{
		Token: token,
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
	return u.UserRepo.Save(user)
}

func (u *UserApplication) QrCode(dto QrCodeDTO) (*QrCodeRet, error) {
	qrCodeValue := cache.Get(dto.Ticket).(cache.QrCodeValue)
	if qrCodeValue.Status == QrCodeStatusAuthorized && qrCodeValue.Token == "" {
		parseJwt, err := jwtutils.ParseJwt(qrCodeValue.TemporaryToken, consts.JwtSecret)
		if err != nil {
			return nil, err
		}
		claims := jwtutils.Claims{
			RegisteredClaims: jwt.RegisteredClaims{
				Issuer:    consts.SystemName,
				Subject:   consts.JwtSubject,
				IssuedAt:  jwt.NewNumericDate(time.Now()),
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 12)),
			},
			User: jwtutils.User{
				Id:       parseJwt.User.Id,
				Username: parseJwt.User.Username,
			},
		}
		token, err := jwtutils.GenerateJwt(claims, consts.JwtSecret)
		if err != nil {
			return nil, err
		}
		qrCodeValue.Token = token
		qrCodeValue.TemporaryToken = ""
		cache.Save(dto.Ticket, qrCodeValue)
	}
	return &QrCodeRet{
		Ticket: dto.Ticket,
		Status: qrCodeValue.Status,
		Token:  qrCodeValue.Token,
	}, nil
}

func (u *UserApplication) ConfirmLogin(dto ConfirmLoginDTO) error {
	qrCodeValue := cache.Get(dto.Ticket).(cache.QrCodeValue)
	if qrCodeValue.Status != QrCodeStatusUnauthorized {
		return errors.New("二维码已被扫描")
	}
	qrCodeValue.Status = QrCodeStatusAuthorized
	qrCodeValue.TemporaryToken = dto.TemporaryToken
	cache.Save(dto.Ticket, qrCodeValue)
	return nil
}

var _ IUserApplication = &UserApplication{}
