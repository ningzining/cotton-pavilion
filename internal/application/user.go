package application

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
	"user-center/internal/application/types"
	"user-center/internal/domain/entity"
	"user-center/internal/domain/repository"
	"user-center/internal/domain/service"
	"user-center/internal/infrastructure/consts"
	"user-center/internal/infrastructure/util/cryptoutil"
	"user-center/internal/infrastructure/util/jwtutil"
	"user-center/pkg/code"
	"user-center/pkg/errors"
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

func (u *UserApplication) Login(dto types.LoginDTO) (*types.LoginRet, error) {
	user, err := u.UserRepo.FindByMobile(dto.Mobile)
	if err != nil {
		return nil, errors.WithCode(code.ErrDatabase, err.Error())
	}
	if user.Password != cryptoutil.Md5Password(dto.Mobile, dto.Password) {
		return nil, errors.WithCode(code.ErrPasswordIncorrect, "密码不正确")
	}

	claims := jwtutil.Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:   consts.SystemName,
			Subject:  consts.JwtSubject,
			IssuedAt: jwt.NewNumericDate(time.Now()),
		},
		User: jwtutil.User{
			Id:       user.ID,
			Username: user.Username,
		},
	}
	token, err := jwtutil.GenerateJwt(claims, consts.JwtSecret)
	if err != nil {
		return nil, errors.WithCode(code.ErrTokenGenerate, err.Error())
	}

	return &types.LoginRet{
		Token: token,
	}, nil
}

func (u *UserApplication) Register(dto types.RegisterDTO) error {
	user := &entity.User{
		Username: dto.Username,
		Mobile:   dto.Mobile,
		Email:    dto.Email,
		Password: cryptoutil.Md5Password(dto.Mobile, dto.Password),
	}
	if err := u.UserRepo.Save(user); err != nil {
		return errors.WithCode(code.ErrDatabase, err.Error())
	}
	return nil
}

func (u *UserApplication) QrCode(dto types.QrCodeDTO) (*types.QrCodeRet, error) {
	ticket := u.QrCodeService.GetTicket(dto.Conn)
	qrCode, err := u.QrCodeService.GetQrCode(ticket)
	if err != nil {
		return nil, errors.WithCode(code.ErrQrCodeInvalid, err.Error())
	}

	// 如果已经授权，则在缓存中移除相关信息
	if qrCode.IsAuthorized() {
		u.QrCodeService.Remove(dto.Conn, qrCode.Ticket)
	}

	return &types.QrCodeRet{
		Ticket: qrCode.Ticket,
		Status: qrCode.Status,
		Token:  qrCode.Token,
	}, nil
}

func (u *UserApplication) ScanQrCode(dto types.ScanQrCodeDTO) (*types.ScanQrCodeRet, error) {
	qrCode, err := u.QrCodeService.GetQrCode(dto.Ticket)
	if err != nil {
		return nil, errors.WithCode(code.ErrQrCodeInvalid, err.Error())
	}
	if qrCode.IsExpired() {
		return nil, errors.WithCode(code.ErrQrCodeExpired, "二维码已过期")
	}
	if !qrCode.IsUnauthorized() {
		return nil, errors.WithCode(code.ErrQrCodeExpired, "二维码已被扫描")
	}

	// 生成临时token
	claims := jwtutil.Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    consts.SystemName,
			Subject:   consts.JwtTemporarySubject,
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 10)),
		},
	}
	temporaryToken, err := jwtutil.GenerateJwt(claims, consts.JwtSecret)
	if err != nil {
		return nil, err
	}

	// 更新二维码为授权中状态
	qrCode.UpdateAuthorizing(temporaryToken)
	u.QrCodeService.SaveQrCode(qrCode)
	return &types.ScanQrCodeRet{
		TemporaryToken: temporaryToken,
	}, nil
}

func (u *UserApplication) ConfirmLogin(dto types.ConfirmLoginDTO) error {
	qrCode, err := u.QrCodeService.GetQrCode(dto.Ticket)
	if err != nil {
		return errors.WithCode(code.ErrQrCodeInvalid, err.Error())
	}
	if !qrCode.IsAuthorizing() {
		return errors.WithCode(code.ErrQrCodeInvalid, "二维码未扫描")
	}
	if dto.TemporaryToken != qrCode.TemporaryToken {
		return errors.WithCode(code.ErrQrCodeInvalid, "扫描二维码和确认的手机不一致")
	}
	parseJwt, err := jwtutil.ParseJwt(dto.Token, consts.JwtSecret)
	if err != nil {
		return errors.WithCode(code.ErrTokenInvalid, err.Error())
	}

	// todo: 按需获取用户信息并放入token当中
	// 生成web端的token
	claims := jwtutil.Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    consts.SystemName,
			Subject:   consts.JwtSubject,
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 12)),
		},
		User: jwtutil.User{
			Id:       parseJwt.User.Id,
			Username: parseJwt.User.Username,
		},
	}
	token, err := jwtutil.GenerateJwt(claims, consts.JwtSecret)
	if err != nil {
		return errors.WithCode(code.ErrTokenGenerate, err.Error())
	}

	// 更新二维码状态为已授权
	qrCode.UpdateAuthorized(token)
	u.QrCodeService.SaveQrCode(qrCode)
	return nil
}
