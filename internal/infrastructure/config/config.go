package config

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"user-center/pkg/logger"
)

var Config *viper.Viper

func init() {
	Config = viper.New()
	Config.SetConfigName("config")
	Config.SetConfigType("yaml")
	Config.AddConfigPath("./config")
	Config.AddConfigPath("../config")
	Config.AddConfigPath("../../config")
	Config.AddConfigPath("../../../config")
	if err := Config.ReadInConfig(); err != nil {
		logger.Fatal("配置文件读取失败", zap.String("error", err.Error()))
	}
}
