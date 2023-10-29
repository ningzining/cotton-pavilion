package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
	application "user-center/internal/application/impl"
	"user-center/internal/infrastructure/db/mysql"
	"user-center/internal/infrastructure/logger"
	"user-center/internal/infrastructure/persistence"
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

	// 初始化repo层
	repositories := persistence.NewRepositories(mysql.DB())
	// 初始化应用层
	app := application.New(repositories)
	// 初始化接口层
	userHandler := interfaces.NewUser(app)

	// 注册路由
	v1Group := engine.Group("/v1")
	v1Group.POST("/user/register", userHandler.Register)
	v1Group.POST("/user/login", userHandler.Login)

	// 启动http服务器
	if err := engine.Run(":8080"); err != nil {
		logger.Info("http server start error", zap.String("err", err.Error()))
		return
	}
}
