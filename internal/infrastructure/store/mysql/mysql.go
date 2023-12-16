package mysql

import (
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sync"
	"time"
	"user-center/internal/domain/model"
	"user-center/internal/domain/repository"
	"user-center/internal/infrastructure/store"
	"user-center/pkg/logger"
)

type Repository struct {
	DB *gorm.DB
}

func (r *Repository) UserRepository() repository.UserRepository {
	return NewUserRepository(r)
}

var (
	mysqlFactory store.Factory
	once         sync.Once
)

func GetMysqlFactory(dsn string) store.Factory {
	var err error
	var dbIns *gorm.DB
	once.Do(func() {
		dbIns, err = newMysql(dsn)
		mysqlFactory = &Repository{DB: dbIns}
	})

	if mysqlFactory == nil || err != nil {
		logger.Fatal("mysql连接异常", zap.String("error", err.Error()))
		return nil
	}
	return mysqlFactory
}

func newMysql(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxIdleConns(100)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Minute)

	logger.Info("mysql start success")
	return db, nil
}

func (r *Repository) AutoMigrate() {
	_ = r.DB.AutoMigrate(&model.User{})
}

func (r *Repository) Clean() {
	_ = r.DB.Migrator().DropTable(&model.User{})
}
