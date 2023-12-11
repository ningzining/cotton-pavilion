package main

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"user-center/internal/infrastructure/config"
	"user-center/internal/interfaces/middleware"
	"user-center/internal/interfaces/router"
	"user-center/pkg/logger"
)

func main() {
	engine := gin.New()
	// 跨域处理
	engine.Use(gin.Recovery())
	engine.Use(middleware.Logger())
	engine.Use(middleware.RequestID())
	engine.Use(middleware.Cors())

	// 注册路由
	router.Register(engine)

	// 启动http服务器
	if err := engine.Run(config.Config.GetString("port")); err != nil {
		logger.Info("http server start error", zap.String("err", err.Error()))
		return
	}
}
