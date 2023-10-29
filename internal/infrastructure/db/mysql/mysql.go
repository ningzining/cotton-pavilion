package mysql

import (
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
	"user-center/internal/infrastructure/logger"
)

var _db *gorm.DB

func init() {
	dsn := "root:root@tcp(127.0.0.1:3306)/user-center?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Error("mysql start error", zap.String("error", err.Error()))
		return
	}
	_db = db
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(500)
	sqlDB.SetConnMaxLifetime(time.Hour)
	logger.Info("mysql start success")
}

func DB() *gorm.DB {
	if _db != nil {
		return _db
	}
	return nil
}
