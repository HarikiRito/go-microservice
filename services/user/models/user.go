package models

import (
	"gorm.io/gorm"
)

type UserEntity struct {
	gorm.Model
	Name     string
	Age      uint
	Password string
}

func (UserEntity) TableName() string {
	return "users"
}
