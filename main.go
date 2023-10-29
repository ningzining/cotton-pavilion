package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
	"user-center/infrastructure/logger"
)

func main() {
	engine := gin.Default()
	engine.Use(cors.New(
		cors.Config{
			AllowAllOrigins:  true,
			AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
			AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type"},
			AllowCredentials: false,
			MaxAge:           12 * time.Hour,
		},
	))

	if err := engine.Run(":8080"); err != nil {
		logger.Info("http server start error", zap.String("err", err.Error()))
		return
	}
}
