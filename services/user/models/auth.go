package models

import "gorm.io/gorm"

type AuthEntity struct {
	gorm.Model
	UserId       string
	AccessToken  string
	RefreshToken string
}

func (AuthEntity) TableName() string {
	return "auth_tokens"
}
