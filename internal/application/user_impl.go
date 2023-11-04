package application

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
	"user-center/internal/domain/entity"
	"user-center/internal/domain/repository"
	"user-center/internal/domain/service"
	"user-center/internal/infrastructure/consts"
	"user-center/internal/infrastructure/utils/crypto"
	"user-center/internal/infrastructure/utils/jwtutils"
)

type UserApplication struct {
	UserRepo      repository.IUserRepository
	QrCodeService service.IQrCodeService
}

func NewUserApplication(userRepository repository.IUserRepository, codeService service.IQrCodeService) *UserApplication {
	return &UserApplication{
		UserRepo:      userRepository,
		QrCodeService: codeService,
	}
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

func (u *UserApplication) QrCode(dto QrCodeDTO) *QrCodeRet {
	ticket := u.QrCodeService.GetTicket(dto.Conn)
	qrCode := u.QrCodeService.GetQrCode(ticket)
	u.QrCodeService.SaveConn(dto.Conn, qrCode.Ticket)

	// 如果已经授权，则在缓存中移除相关信息
	if qrCode.IsAuthorized() {
		u.QrCodeService.RemoveTicket(qrCode.Ticket)
		u.QrCodeService.RemoveConn(dto.Conn)
	}

	return &QrCodeRet{
		Ticket: qrCode.Ticket,
		Status: qrCode.Status,
		Token:  qrCode.Token,
	}
}

func (u *UserApplication) ScanQrCode(dto ScanQrCodeDTO) (*ScanQrCodeRet, error) {
	qrCode := u.QrCodeService.GetQrCode(dto.Ticket)
	if !qrCode.IsUnauthorized() {
		return nil, errors.New("二维码已扫描")
	}

	// 生成临时token
	claims := jwtutils.Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    consts.SystemName,
			Subject:   consts.JwtTemporarySubject,
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 10)),
		},
	}
	temporaryToken, err := jwtutils.GenerateJwt(claims, consts.JwtSecret)
	if err != nil {
		return nil, err
	}

	// 更新二维码为授权中状态
	qrCode.UpdateAuthorizing(temporaryToken)
	u.QrCodeService.SaveQrCode(qrCode)
	return &ScanQrCodeRet{
		TemporaryToken: temporaryToken,
	}, nil
}

func (u *UserApplication) ConfirmLogin(dto ConfirmLoginDTO) error {
	qrCode := u.QrCodeService.GetQrCode(dto.Ticket)
	if !qrCode.IsAuthorizing() {
		return errors.New("二维码已授权,请刷新")
	}
	if dto.TemporaryToken != qrCode.TemporaryToken {
		return errors.New("二维码已过期,请刷新")
	}
	parseJwt, err := jwtutils.ParseJwt(dto.Token, consts.JwtSecret)
	if err != nil {
		return err
	}

	// todo: 按需获取用户信息并放入token当中
	// 生成web端的token
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
		return err
	}

	// 更新二维码状态为已授权
	qrCode.UpdateAuthorized(token)
	u.QrCodeService.SaveQrCode(qrCode)
	return nil
}

var _ IUserApplication = &UserApplication{}
