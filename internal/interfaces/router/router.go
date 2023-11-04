package router

import (
	"github.com/gin-gonic/gin"
	"user-center/internal/application"
	"user-center/internal/infrastructure/persistence"
	"user-center/internal/interfaces"
)

func Register(engine *gin.Engine) {
	// 初始化repo层
	repositories := persistence.NewRepositories()
	repositories.AutoMigrate()

	v1Group := engine.Group("/v1")
	registerLoginRouter(v1Group, repositories)
}

func registerLoginRouter(group *gin.RouterGroup, repositories *persistence.Repositories) {
	// 初始化应用层
	app := application.New(repositories)
	// 初始化接口层
	userHandler := interfaces.NewUser(app)
	group.POST("/register", userHandler.Register)
	group.POST("/login", userHandler.Login)
	group.GET("/qr-code", userHandler.QrCode)
	group.POST("/confirm-login", userHandler.ConfirmLogin)
}
