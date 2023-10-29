package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
	application "user-center/internal/application/impl"
	repoistory "user-center/internal/domain/repository/impl"
	"user-center/internal/infrastructure/logger"
	"user-center/internal/interfaces"
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

	// 初始化依赖
	repo := repoistory.New()
	app := application.New(repo)
	userHandler := interfaces.NewUser(app.IUserApplication)

	// 注册路由
	v1Group := engine.Group("/v1")
	v1Group.POST("/user/register", userHandler.Register)

	// 启动http服务器
	if err := engine.Run(":8080"); err != nil {
		logger.Info("http server start error", zap.String("err", err.Error()))
		return
	}
}
