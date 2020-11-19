package database

import "user/models"

func MigrateTable() {
	db := GetDBInstance()
	err := db.AutoMigrate(&models.UserEntity{}, &models.AuthEntity{})
	if err != nil {
		panic("Migrate table error")
	}
}
