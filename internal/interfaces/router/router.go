package router

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"user-center/internal/application"
	"user-center/internal/infrastructure/cache/redis_cache"
	"user-center/internal/infrastructure/persistence"
	"user-center/internal/infrastructure/service"
	"user-center/internal/interfaces"
	"user-center/internal/interfaces/middleware"
)

func Register(engine *gin.Engine) {
	// 初始化redis
	// todo: 待使用
	redis_cache.NewRedisCache(viper.GetString("redis.addr"), viper.GetString("redis.password"), viper.GetInt("redis.db"))
	// 初始化repo层
	repositories := persistence.NewRepositories(viper.GetString("mysql.dsn"))
	repositories.AutoMigrate()
	// 初始化service层
	services := service.New(repositories)
	// 初始化应用层
	userApplication := application.NewUserApplication(repositories.UserRepository, services.QrCodeService)

	v1Group := engine.Group("/v1")
	registerLoginRouter(v1Group, userApplication)
}

func registerLoginRouter(group *gin.RouterGroup, app *application.UserApplication) {
	// 初始化接口层
	userHandler := interfaces.NewUserHandler(app)
	group.POST("/register", userHandler.Register)
	group.POST("/login", userHandler.Login)

	group.GET("/qr-code", userHandler.QrCode)
	group.GET("/scan-qr-code", middleware.JwtMiddleware(), userHandler.ScanQrCode)
	group.GET("/confirm-login", middleware.JwtMiddleware(), userHandler.ConfirmLogin)
}
