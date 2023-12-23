package server

import (
	"github.com/gin-gonic/gin"
	"github.com/ningzining/cotton-pavilion/internal/infrastructure/cache/redis_cache"
	"github.com/ningzining/cotton-pavilion/internal/infrastructure/config"
	"github.com/ningzining/cotton-pavilion/internal/infrastructure/store"
	"github.com/ningzining/cotton-pavilion/internal/infrastructure/store/mysql"
	"github.com/ningzining/cotton-pavilion/internal/interfaces/middleware"
	"github.com/ningzining/cotton-pavilion/pkg/logger"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type Server struct {
	Engine *gin.Engine
}

func New() *Server {
	initConfig()

	server := createServer()

	return server
}

func (s Server) Run() {
	// 启动http服务器
	if err := s.Engine.Run(viper.GetString("port")); err != nil {
		logger.Fatal("http server start error", zap.String("err", err.Error()))
		return
	}
}

func initConfig() {
	// 读取配置文件
	config.LoadConfig()
	// 初始化redis
	_ = redis_cache.NewRedisCache(viper.GetString("redis.addr"), viper.GetString("redis.password"), viper.GetInt("redis.db"))
	// 初始化repo层
	repositoryFactory := mysql.GetMysqlFactory(viper.GetString("mysql.dsn"))
	repositoryFactory.AutoMigrate()
	store.SetClient(repositoryFactory)
}

func createServer() *Server {
	server := &Server{
		Engine: gin.New(),
	}

	initServer(server)
	return server
}

func initServer(s *Server) {
	s.installMiddlewares()
}

func (s Server) installMiddlewares() {
	s.Engine.Use(middleware.DefaultMiddlewares()...)
}
