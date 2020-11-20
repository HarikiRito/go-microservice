package repositories

import (
	"user/models"
	"user/services"
)

type UserEntity = models.UserEntity
type UserRepository struct {
}

var commonRepo = NewCommonRepo()
var authService = services.AuthService{}

func (r *UserRepository) GetAll() []UserEntity {
	var records []UserEntity
	commonRepo.GetAll(&records)
	return records
}

func (r *UserRepository) FindById(id uint) UserEntity {
	var record UserEntity
	commonRepo.FindById(&record, id)
	return record
}

func (r *UserRepository) Create(model UserEntity) UserEntity {
	newPassword, _ := authService.HashPassword(model.Password)
	model.Password = newPassword
	commonRepo.Create(&model)
	return model
}
