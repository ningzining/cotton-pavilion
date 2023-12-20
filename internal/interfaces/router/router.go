package router

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"user-center/internal/domain/service"
	"user-center/internal/infrastructure/store/mysql"
	"user-center/internal/infrastructure/third_party"
	"user-center/internal/interfaces"
	"user-center/internal/interfaces/middleware"
)

func Register(engine *gin.Engine) {
	storeIns := mysql.GetMysqlFactory(viper.GetString("mysql.dsn"))
	svc := service.NewService(storeIns)
	ossConfig := third_party.OssConfig{
		Bucket:          viper.GetString("oss.Bucket"),
		Endpoint:        viper.GetString("oss.Endpoint"),
		AccessKeyID:     viper.GetString("oss.AccessKeyId"),
		AccessKeySecret: viper.GetString("oss.AccessKeySecret"),
	}
	oss, err := third_party.NewOssClient(ossConfig)
	if err != nil {
		panic("oss服务初始化失败")
	}

	v1Group := engine.Group("/v1")
	v1Group.Use(middleware.JwtMiddleware())
	{
		// 初始化接口层
		userHandler := interfaces.NewUserHandler(storeIns, svc, oss)
		engine.POST("/register", userHandler.Register)
		engine.POST("/login", userHandler.Login)
		engine.GET("/qr-code", userHandler.QrCode)
	}
	{
		// 初始化接口层
		userHandler := interfaces.NewUserHandler(storeIns, svc, oss)

		v1Group.GET("/scan-qr-code", userHandler.ScanQrCode)
		v1Group.GET("/confirm-login", userHandler.ConfirmLogin)
	}
	{
		imageHandler := interfaces.NewImageHandler(storeIns, svc, oss)
		commonGroup := v1Group.Group("/common")
		commonGroup.POST("/upload", imageHandler.Upload)
	}
}
