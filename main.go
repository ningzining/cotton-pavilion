package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"user-center/internal/infrastructure/config"
	"user-center/internal/interfaces/middleware"
	"user-center/internal/interfaces/router"
	"user-center/pkg/logger"
)

func main() {
	config.LoadConfig()

	engine := gin.New()
	// 跨域处理
	engine.Use(gin.Recovery())
	engine.Use(middleware.Cors())
	engine.Use(middleware.RequestID())
	engine.Use(middleware.Logger())

	// 注册路由
	router.Register(engine)

	// 启动http服务器
	if err := engine.Run(viper.GetString("port")); err != nil {
		logger.Info("http server start error", zap.String("err", err.Error()))
		return
	}
}
