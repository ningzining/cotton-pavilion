package config

import (
	"github.com/spf13/viper"
	"testing"
)

func TestInit(t *testing.T) {
	mysql := viper.GetString("mysql.dsn")
	if mysql == "" {
		t.Errorf("配置文件读取失败")
	}
	t.Logf("%s", mysql)
}
