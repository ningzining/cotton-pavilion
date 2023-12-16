package main

import (
	"user-center/internal/infrastructure/config"
	"user-center/internal/infrastructure/server"
	"user-center/internal/interfaces/router"
)

func main() {
	config.LoadConfig()

	// 初始化对象
	s := server.New()
	// 注册路由
	router.Register(s.Engine)
	s.Run()
}
