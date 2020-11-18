package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"sync"
)

var lock = &sync.Mutex{}
var db *gorm.DB

func connectDatabase() {
	host := "127.0.0.1"
	user := "postgres"
	password := "postgres"
	port := "54321"
	dbname := "sayang"
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", host, user, password, dbname, port)
	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db = conn
}

func GetDBInstance() *gorm.DB {
	if db == nil {
		lock.Lock()
		// Lock so only one goroutine at a time can access
		defer lock.Unlock()
		connectDatabase()
	}
	return db
}
