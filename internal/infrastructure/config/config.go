package config

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"user-center/internal/infrastructure/logger"
)

type Config struct {
}

func Init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	viper.AddConfigPath("../config")
	viper.AddConfigPath("../../config")
	viper.AddConfigPath("../../../config")
	err := viper.ReadInConfig()
	if err != nil {
		logger.Fatal("配置文件读取失败", zap.String("error", err.Error()))
	}
}
