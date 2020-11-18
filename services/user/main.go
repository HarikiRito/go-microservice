package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"user/database"
	"user/models"
	"user/repositories"
)

type User = models.User

func main() {
	db := database.GetInstance()
	db.AutoMigrate(&User{})
	var users []User
	var user User
	common := repositories.CommonRepository{Db: db}
	//spew.Dump(a)
	common.GetAll(&users)
	common.FindById(&user, 2)
	//spew.Dump(users)

	r := gin.Default()
	r.GET("/users", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"ok": users,
		})
	})
	r.Run()
}
