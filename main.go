package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
	"user-center/internal/infrastructure/config"
	"user-center/internal/infrastructure/logger"
	"user-center/internal/interfaces/router"
)

func main() {
	engine := gin.Default()
	// 跨域处理
	engine.Use(cors.New(
		cors.Config{
			AllowAllOrigins:  true,
			AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
			AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type"},
			AllowCredentials: false,
			MaxAge:           12 * time.Hour,
		},
	))

	// 注册路由
	router.Register(engine)

	// 启动http服务器
	if err := engine.Run(config.Config.GetString("port")); err != nil {
		logger.Info("http server start error", zap.String("err", err.Error()))
		return
	}
}
