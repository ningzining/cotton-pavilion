package router

import (
	"github.com/gin-gonic/gin"
	"user-center/internal/application"
	"user-center/internal/domain/service"
	"user-center/internal/infrastructure/cache/redis_cache"
	"user-center/internal/infrastructure/persistence"
	"user-center/internal/interfaces"
	"user-center/internal/interfaces/middleware"
)

func Register(engine *gin.Engine) {
	// todo: 待调整 repo层，注入cache和db
	// 初始化cache缓存层
	redisCache := redis_cache.NewRedisCache()
	// 初始化repo层
	repositories := persistence.NewRepositories()
	repositories.AutoMigrate()
	// 初始化service层
	services := service.New()
	// 初始化应用层
	app := application.New(repositories, services, redisCache)

	v1Group := engine.Group("/v1")
	registerLoginRouter(v1Group, app)
}

func registerLoginRouter(group *gin.RouterGroup, app *application.Application) {
	// 初始化接口层
	userHandler := interfaces.NewUser(app)
	group.POST("/register", userHandler.Register)
	group.POST("/login", userHandler.Login)

	group.GET("/qr-code", userHandler.QrCode)
	group.GET("/scan-qr-code", middleware.JwtMiddleware(), userHandler.ScanQrCode)
	group.GET("/confirm-login", middleware.JwtMiddleware(), userHandler.ConfirmLogin)
}
