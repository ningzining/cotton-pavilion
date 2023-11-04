package router

import (
	"github.com/gin-gonic/gin"
	"user-center/internal/application"
	"user-center/internal/domain/service"
	"user-center/internal/infrastructure/persistence"
	"user-center/internal/interfaces"
	"user-center/internal/interfaces/middleware"
)

func Register(engine *gin.Engine) {
	// 初始化repo层
	repositories := persistence.NewRepositories()
	repositories.AutoMigrate()
	// 初始化service层
	services := service.New()

	v1Group := engine.Group("/v1")
	registerLoginRouter(v1Group, repositories, services)
}

func registerLoginRouter(group *gin.RouterGroup, repositories *persistence.Repositories, services *service.Services) {
	// 初始化应用层
	app := application.New(repositories, services)
	// 初始化接口层
	userHandler := interfaces.NewUser(app)
	group.POST("/register", userHandler.Register)
	group.POST("/login", userHandler.Login)

	group.GET("/qr-code", userHandler.QrCode)
	group.GET("/scan-qr-code", middleware.JwtMiddleware(), userHandler.ScanQrCode)
	group.GET("/confirm-login", middleware.JwtMiddleware(), userHandler.ConfirmLogin)
}
