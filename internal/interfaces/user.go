package interfaces

import (
	"github.com/gin-gonic/gin"
	"time"
	"user-center/internal/application/types"
	"user-center/internal/domain/entity/enum"
	"user-center/internal/infrastructure/application"
	"user-center/internal/infrastructure/logger"
	"user-center/internal/infrastructure/utils/wsutils"
	"user-center/pkg/code"
	"user-center/pkg/errors"
	"user-center/pkg/response"
)

type UserHandler struct {
}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

// Register
//
//	@Summary	注册
//	@Tags		login
//	@Param		req	body		application.RegisterDTO	true	"req"
//	@Success	200	{object}	response.Result
//	@Router		/register [post]
func (u *UserHandler) Register(ctx *gin.Context) {
	var dto types.RegisterDTO
	if err := ctx.ShouldBind(&dto); err != nil {
		response.Error(ctx, errors.WithCode(code.ErrBind, err.Error()))
		return
	}

	if err := application.UserApplication().Register(dto); err != nil {
		response.Error(ctx, err)
		return
	}
	response.Success(ctx, nil)
}

// Login
//
//	@Summary	登录
//	@Tags		login
//	@Param		req	body		application.LoginDTO	true	"req"
//	@Success	200	{object}	response.Result{data=application.LoginRet}
//	@Router		/login [post]
func (u *UserHandler) Login(ctx *gin.Context) {
	var dto types.LoginDTO
	if err := ctx.ShouldBind(&dto); err != nil {
		response.Error(ctx, errors.WithCode(code.ErrBind, err.Error()))
		return
	}

	data, err := application.UserApplication().Login(dto)
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
	dto := types.QrCodeDTO{
		Conn: conn,
	}
	defer conn.Close()
	for {
		codeRet := application.UserApplication().QrCode(dto)
		time.Sleep(time.Second)
		_ = conn.WriteJSON(codeRet)
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
//	@Param		req	query		application.ScanQrCodeDTO	true	"req"
//	@Success	200	{object}	response.Result
//	@Router		/scan-qr-code [get]
func (u *UserHandler) ScanQrCode(ctx *gin.Context) {
	var dto types.ScanQrCodeDTO
	if err := ctx.ShouldBindQuery(&dto); err != nil {
		response.Error(ctx, errors.WithCode(code.ErrBind, err.Error()))
		return
	}
	if err := ctx.ShouldBindHeader(&dto); err != nil {
		response.Error(ctx, errors.WithCode(code.ErrMissingHeader, err.Error()))
		return
	}

	data, err := application.UserApplication().ScanQrCode(dto)
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
	var dto types.ConfirmLoginDTO
	if err := ctx.ShouldBind(&dto); err != nil {
		response.Error(ctx, errors.WithCode(code.ErrBind, err.Error()))
		return
	}
	if err := ctx.ShouldBindHeader(&dto); err != nil {
		response.Error(ctx, errors.WithCode(code.ErrMissingHeader, err.Error()))
		return
	}

	if err := application.UserApplication().ConfirmLogin(dto); err != nil {
		response.Error(ctx, err)
		return
	}
	response.Success(ctx, nil)
}
