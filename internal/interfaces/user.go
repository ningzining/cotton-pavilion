package interfaces

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"net/http"
	"time"
	"user-center/internal/application"
	"user-center/internal/infrastructure/cache"
	"user-center/internal/infrastructure/logger"
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

func (u *UserHandler) QrCode(ctx *gin.Context) {
	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		logger.Error(err.Error())
		return
	}
	dto := application.QrCodeDTO{
		Conn:   conn,
		Ticket: uuid.New().String(),
	}
	cache.Save(dto.Ticket, application.QrCodeStatusUnauthorized)
	defer conn.Close()
	for {
		codeRet := u.UserApp.QrCode(dto)
		time.Sleep(time.Second)
		if err != nil {
			logger.Error(err.Error())
			return
		}
		conn.WriteJSON(codeRet)
	}
}
