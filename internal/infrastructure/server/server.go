package server

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"user-center/internal/infrastructure/cache/redis_cache"
	"user-center/internal/infrastructure/config"
	"user-center/internal/infrastructure/store"
	"user-center/internal/infrastructure/store/mysql"
	"user-center/internal/interfaces/middleware"
	"user-center/pkg/logger"
)

type Server struct {
	Engine *gin.Engine
}

func New() *Server {
	// 读取配置文件
	config.LoadConfig()
	// 初始化redis
	_ = redis_cache.NewRedisCache(viper.GetString("redis.addr"), viper.GetString("redis.password"), viper.GetInt("redis.db"))
	// 初始化repo层
	repositoryFactory := mysql.GetMysqlFactory(viper.GetString("mysql.dsn"))
	repositoryFactory.AutoMigrate()
	store.SetClient(repositoryFactory)

	engine := gin.New()
	engine.Use(middleware.DefaultMiddlewares()...)

	return &Server{
		Engine: engine,
	}
}

func (s Server) Run() {
	// 启动http服务器
	if err := s.Engine.Run(viper.GetString("port")); err != nil {
		logger.Fatal("http server start error", zap.String("err", err.Error()))
		return
	}
}
