package repositories

import (
	"gorm.io/gorm"
)

type CommonRepository struct {
	Db *gorm.DB
}

func (r CommonRepository) Create(model interface{}) *gorm.DB {
	return r.Db.Create(model)
}

func (r CommonRepository) GetAll(records interface{}) *gorm.DB {
	return r.Db.Find(records)
}

func (r CommonRepository) Update(model interface{}) *gorm.DB {
	return r.Db.Save(model)
}

func (r CommonRepository) FindById(model interface{}, id uint) *gorm.DB {
	return r.Db.First(model, id)
}

func (r CommonRepository) DeleteById(model interface{}, id uint) {
	r.Db.Delete(model, id)
}
