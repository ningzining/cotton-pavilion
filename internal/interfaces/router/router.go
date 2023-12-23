package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ningzining/cotton-pavilion/internal/domain/service"
	"github.com/ningzining/cotton-pavilion/internal/infrastructure/store/mysql"
	"github.com/ningzining/cotton-pavilion/internal/interfaces"
	"github.com/ningzining/cotton-pavilion/internal/interfaces/middleware"
	"github.com/spf13/viper"
)

func Register(engine *gin.Engine) {
	storeIns := mysql.GetMysqlFactory(viper.GetString("mysql.dsn"))
	svc := service.NewService(storeIns)

	v1Group := engine.Group("/v1")
	v1Group.Use(middleware.JwtMiddleware())
	{
		// 初始化接口层
		userHandler := interfaces.NewUserHandler(storeIns, svc)
		engine.POST("/register", userHandler.Register)
		engine.POST("/login", userHandler.Login)
		engine.GET("/qr-code", userHandler.QrCode)
	}
	{
		// 初始化接口层
		userHandler := interfaces.NewUserHandler(storeIns, svc)

		v1Group.GET("/scan-qr-code", userHandler.ScanQrCode)
		v1Group.GET("/confirm-login", userHandler.ConfirmLogin)
	}
	{
		imageHandler := interfaces.NewImageHandler(storeIns, svc)
		commonGroup := v1Group.Group("/common")
		commonGroup.POST("/upload", imageHandler.Upload)
	}
}
