package main

import (
	"github.com/ningzining/cotton-pavilion/internal/infrastructure/server"
	"github.com/ningzining/cotton-pavilion/internal/interfaces/router"
)

func main() {
	// 初始化对象
	s := server.New()
	// 注册路由
	router.Register(s.Engine)

	s.Run()
}
