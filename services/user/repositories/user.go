package repositories

import "user/models"

type User = models.User
type UserRepository struct {
}

func (r *UserRepository) GetAll() []User {
	var records []User
	GetCommonRepo().GetAll(&records)
	return records
}

func (r *UserRepository) FindById(id uint) User {
	var record User
	GetCommonRepo().FindById(&record, id)
	return record
}
