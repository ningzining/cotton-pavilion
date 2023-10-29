package entity

import (
	"user-center/internal/infrastructure/db/mysql"
)

type User struct {
	BaseModel
	Username string
	Mobile   string
	Email    string
	Password string
}

func init() {
	mysql.DB().AutoMigrate(User{})
}
