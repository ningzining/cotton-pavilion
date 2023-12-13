package application

import (
	"github.com/spf13/viper"
	"user-center/internal/application"
	"user-center/internal/infrastructure/cache/redis_cache"
	"user-center/internal/infrastructure/persistence"
	"user-center/internal/infrastructure/service"
)

type Application struct {
	UserApplication *application.UserApplication
}

var std *Application

func newApplication() *Application {
	// 初始化redis,todo: 待使用
	_ = redis_cache.NewRedisCache(viper.GetString("redis.addr"), viper.GetString("redis.password"), viper.GetInt("redis.db"))
	// 初始化repo层
	repositories := persistence.NewRepositories(viper.GetString("mysql.dsn"))
	repositories.AutoMigrate()
	// 初始化service层
	services := service.New()
	// 初始化应用层
	return &Application{
		UserApplication: application.NewUserApplication(repositories.UserRepository, services.QrCodeService),
	}
}

func UserApplication() *application.UserApplication {
	if std == nil {
		std = newApplication()
	}
	return std.UserApplication
}
