package interfaces

import (
	"github.com/gin-gonic/gin"
	"user-center/internal/application"
	"user-center/pkg/response"
)

func NewUser(app *application.Application) *UserHandler {
	return &UserHandler{
		UserApp: app.IUserApplication,
	}
}

type UserHandler struct {
	UserApp application.IUserApplication
}

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
