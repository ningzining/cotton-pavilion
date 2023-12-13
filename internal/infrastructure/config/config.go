package config

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"user-center/pkg/logger"
)

func LoadConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	viper.AddConfigPath("../config")
	viper.AddConfigPath("../../config")
	viper.AddConfigPath("../../../config")
	if err := viper.ReadInConfig(); err != nil {
		logger.Fatal("配置文件读取失败", zap.String("error", err.Error()))
	}
}
