package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string
	Mobile   string
	Email    string
	Password string
}
