package model

import "gorm.io/gorm"

type Image struct {
	gorm.Model
	Url          string
	CreateUserId uint
}
