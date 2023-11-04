package persistence

import (
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
	"user-center/internal/domain/entity"
	"user-center/internal/domain/repository"
	"user-center/internal/infrastructure/logger"
)

type Repositories struct {
	db   *gorm.DB
	User repository.IUserRepository
}

func NewRepositories(dsn string) *Repositories {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Fatal("mysql start error", zap.String("error", err.Error()))
		return nil
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(500)
	sqlDB.SetConnMaxLifetime(time.Hour)
	logger.Info("mysql start success")

	return &Repositories{
		db:   db,
		User: NewUserRepository(db),
	}
}

func (r Repositories) AutoMigrate() {
	r.db.AutoMigrate(entity.User{})
}
