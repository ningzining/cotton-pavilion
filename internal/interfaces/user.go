package interfaces

import (
	"github.com/gin-gonic/gin"
	"time"
	"user-center/internal/application"
	"user-center/internal/domain/entity/enum"
	"user-center/internal/infrastructure/consts"
	"user-center/internal/infrastructure/logger"
	"user-center/internal/infrastructure/utils/wsutils"
	"user-center/pkg/response"
)

type UserHandler struct {
	UserApp *application.UserApplication
}

func NewUserHandler(app *application.UserApplication) *UserHandler {
	return &UserHandler{
		UserApp: app,
	}
}

// Register
//
//	@Summary	注册
//	@Tags		login
//	@Accept		json
//	@Produce	json
//	@Param		req	body		application.RegisterDTO	true	"req"
//	@Success	200	{object}	response.Result
//	@Router		/register [post]
func (u *UserHandler) Register(ctx *gin.Context) {
	var dto application.RegisterDTO
	if err := ctx.ShouldBind(&dto); err != nil {
		response.Error(ctx, err)
		return
	}

	if err := u.UserApp.Register(dto); err != nil {
		response.Error(ctx, err)
		return
	}
	response.Success(ctx, nil)
}

// Login
//
//	@Summary	登录
//	@Tags		login
//	@Accept		json
//	@Produce	json
//	@Param		req	body		application.LoginDTO	true	"req"
//	@Success	200	{object}	response.Result{data=application.LoginRet}
//	@Router		/login [post]
func (u *UserHandler) Login(ctx *gin.Context) {
	var dto application.LoginDTO
	if err := ctx.ShouldBind(&dto); err != nil {
		response.Error(ctx, err)
		return
	}

	data, err := u.UserApp.Login(dto)
	if err != nil {
		response.Error(ctx, err)
		return
	}
	response.Success(ctx, data)
}

func (u *UserHandler) QrCode(ctx *gin.Context) {
	conn, err := wsutils.UpGrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		logger.Error(err.Error())
		return
	}
	dto := application.QrCodeDTO{
		Conn: conn,
	}
	defer conn.Close()
	for {
		codeRet := u.UserApp.QrCode(dto)
		time.Sleep(time.Second)
		conn.WriteJSON(codeRet)
		// 如果已经授权了，那么需要跳出循环，关闭ws连接
		if codeRet.Status == enum.QrCodeStatusAuthorized {
			return
		}
	}
}

// ScanQrCode
//
//	@Summary	扫描二维码
//	@Tags		login
//	@Accept		json
//	@Produce	json
//	@Param		req	query		application.ScanQrCodeDTO	true	"req"
//	@Success	200	{object}	response.Result
//	@Router		/scan-qr-code [get]
func (u *UserHandler) ScanQrCode(ctx *gin.Context) {
	var dto application.ScanQrCodeDTO
	if err := ctx.ShouldBind(&dto); err != nil {
		response.Error(ctx, err)
		return
	}
	dto.Token = ctx.GetString(consts.ContextKeyToken)

	data, err := u.UserApp.ScanQrCode(dto)
	if err != nil {
		response.Error(ctx, err)
		return
	}
	response.Success(ctx, data)
}

// ConfirmLogin
//
//	@Summary	确定登录
//	@Tags		login
//	@Accept		json
//	@Produce	json
//	@Param		req	query		application.ConfirmLoginDTO	true	"req"
//	@Success	200	{object}	response.Result
//	@Router		/confirm-login [get]
func (u *UserHandler) ConfirmLogin(ctx *gin.Context) {
	var dto application.ConfirmLoginDTO
	if err := ctx.ShouldBind(&dto); err != nil {
		response.Error(ctx, err)
		return
	}
	dto.Token = ctx.GetString(consts.ContextKeyToken)

	if err := u.UserApp.ConfirmLogin(dto); err != nil {
		response.Error(ctx, err)
		return
	}
	response.Success(ctx, nil)
}
