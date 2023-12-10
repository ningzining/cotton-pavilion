package application

import (
	"user-center/internal/application"
	"user-center/internal/infrastructure/cache/redis_cache"
	"user-center/internal/infrastructure/config"
	"user-center/internal/infrastructure/persistence"
	"user-center/internal/infrastructure/service"
)

var std = newApplication()

type Application struct {
	UserApplication *application.UserApplication
}

func newApplication() *Application {
	// 初始化redis,todo: 待使用
	redis_cache.NewRedisCache(config.Config.GetString("redis.addr"), config.Config.GetString("redis.password"), config.Config.GetInt("redis.db"))
	// 初始化repo层
	repositories := persistence.NewRepositories(config.Config.GetString("mysql.dsn"))
	repositories.AutoMigrate()
	// 初始化service层
	services := service.New()
	// 初始化应用层
	return &Application{
		UserApplication: application.NewUserApplication(repositories.UserRepository, services.QrCodeService),
	}
}

func UserApplication() *application.UserApplication {
	return std.UserApplication
}
