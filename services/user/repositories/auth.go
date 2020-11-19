package repositories

import "user/models"

type AuthEntity = models.AuthEntity
type AuthRepository struct {
}

func (r *AuthRepository) SaveToken(record AuthEntity) *AuthEntity {
	NewCommonRepo().Create(&record)
	return &record
}
