package router

import (
	"github.com/gin-gonic/gin"
	"user-center/internal/interfaces"
	"user-center/internal/interfaces/middleware"
)

func Register(engine *gin.Engine) {
	v1Group := engine.Group("/v1")
	registerLoginRouter(v1Group)
}

func registerLoginRouter(group *gin.RouterGroup) {
	// 初始化接口层
	userHandler := interfaces.NewUserHandler()
	group.POST("/register", userHandler.Register)
	group.POST("/login", userHandler.Login)

	group.GET("/qr-code", userHandler.QrCode)
	group.GET("/scan-qr-code", middleware.JwtMiddleware(), userHandler.ScanQrCode)
	group.GET("/confirm-login", middleware.JwtMiddleware(), userHandler.ConfirmLogin)
}
