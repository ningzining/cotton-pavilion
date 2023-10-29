package entity

import (
	"gorm.io/gorm"
	"user-center/internal/infrastructure/db/mysql"
)

type User struct {
	gorm.Model
	Username string
	Mobile   string
	Email    string
	Password string
}

func init() {
	mysql.DB().AutoMigrate(User{})
}
