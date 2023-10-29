package interfaces

import (
	"github.com/gin-gonic/gin"
	"user-center/internal/application"
	"user-center/pkg/result"
)

func NewUser(userApp application.IUserApplication) UserHandler {
	return UserHandler{
		UserApp: userApp,
	}
}

type UserHandler struct {
	UserApp application.IUserApplication
}

func (u *UserHandler) Register(ctx *gin.Context) {
	var dto application.RegisterDTO
	if err := ctx.ShouldBind(&dto); err != nil {
		result.Error(ctx, err)
		return
	}
	if err := u.UserApp.Register(dto); err != nil {
		result.Error(ctx, err)
		return
	}
	result.Success(ctx, nil)
}
