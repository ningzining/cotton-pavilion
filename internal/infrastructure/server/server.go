package server

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"user-center/internal/infrastructure/cache/redis_cache"
	"user-center/internal/infrastructure/store"
	"user-center/internal/infrastructure/store/mysql"
	"user-center/pkg/logger"
)

func New() {
	_ = redis_cache.NewRedisCache(viper.GetString("redis.addr"), viper.GetString("redis.password"), viper.GetInt("redis.db"))
	// 初始化repo层
	repositoryFactory, err := mysql.GetMysqlFactory(viper.GetString("mysql.dsn"))
	if err != nil {
		logger.Fatal("mysql start error", zap.String("error", err.Error()))
		return
	}
	repositoryFactory.AutoMigrate()
	store.SetClient(repositoryFactory)
}
