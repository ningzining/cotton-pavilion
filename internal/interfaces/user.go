package interfaces

import (
	"github.com/gin-gonic/gin"
	"time"
	"user-center/internal/application"
	"user-center/internal/application/types"
	"user-center/internal/domain/model/enum"
	"user-center/internal/infrastructure/cache/qr_code_conn_cache"
	"user-center/internal/infrastructure/service"
	"user-center/internal/infrastructure/store"
	"user-center/internal/infrastructure/util/wsutil"
	"user-center/pkg/code"
	"user-center/pkg/errors"
	"user-center/pkg/logger"
	"user-center/pkg/response"
)

type UserHandler struct {
	Application application.Application
}

func NewUserHandler(store store.Factory, service service.Service) *UserHandler {
	return &UserHandler{
		Application: application.NewApplication(store, service),
	}
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

	if err := u.Application.UserApplication().Register(dto); err != nil {
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

	data, err := u.Application.UserApplication().Login(dto)
	if err != nil {
		response.Error(ctx, err)
		return
	}
	response.Success(ctx, data)
}

// QrCode ws pc获取二维码
func (u *UserHandler) QrCode(ctx *gin.Context) {
	conn, err := wsutil.UpGrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		logger.Error(err.Error())
		return
	}
	defer func() {
		_ = conn.Close()
	}()

	dto := types.QrCodeDTO{
		Conn: conn,
	}
	for {
		codeRet, err := u.Application.UserApplication().QrCode(dto)
		if err != nil {
			return
		}
		time.Sleep(time.Second)
		_ = conn.WriteJSON(codeRet)
		// 如果已经授权了，那么需要跳出循环，关闭ws连接
		if codeRet.Status == enum.QrCodeStatusAuthorized {
			return
		}
	}
	qr_code_conn_cache.Remove(conn)
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

	data, err := u.Application.UserApplication().ScanQrCode(dto)
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

	if err := u.Application.UserApplication().ConfirmLogin(dto); err != nil {
		response.Error(ctx, err)
		return
	}
	response.Success(ctx, nil)
}
