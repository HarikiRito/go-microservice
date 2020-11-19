package repositories

import "user/models"

type UserEntity = models.UserEntity
type UserRepository struct {
}

func (r *UserRepository) GetAll() []UserEntity {
	var records []UserEntity
	NewCommonRepo().GetAll(&records)
	return records
}

func (r *UserRepository) FindById(id uint) UserEntity {
	var record UserEntity
	NewCommonRepo().FindById(&record, id)
	return record
}
