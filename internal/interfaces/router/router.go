package router

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"user-center/internal/domain/service"
	"user-center/internal/infrastructure/store/mysql"
	"user-center/internal/interfaces"
	"user-center/internal/interfaces/middleware"
)

func Register(engine *gin.Engine) {
	storeIns := mysql.GetMysqlFactory(viper.GetString("mysql.dsn"))
	svc := service.NewService(storeIns)
	v1Group := engine.Group("/v1")
	{
		// 初始化接口层
		userHandler := interfaces.NewUserHandler(storeIns, svc)

		v1Group.POST("/register", userHandler.Register)
		v1Group.POST("/login", userHandler.Login)

		v1Group.GET("/qr-code", userHandler.QrCode)
		v1Group.GET("/scan-qr-code", middleware.JwtMiddleware(), userHandler.ScanQrCode)
		v1Group.GET("/confirm-login", middleware.JwtMiddleware(), userHandler.ConfirmLogin)
	}
}
