package main

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"user/database"
	"user/models"
	"user/repositories"
	"user/services"
)

type User = models.User

func main() {
	db := database.GetDBInstance()
	db.AutoMigrate(&User{})
	jwtService := services.JwtService{}
	// Sign and get the complete encoded token as a string using the secret
	tokenString, _ := jwtService.GenerateToken(jwt.MapClaims{
		"userId": 1,
	})

	fmt.Println(tokenString)
	userRepo := repositories.UserRepository{}
	users := userRepo.GetAll()
	r := gin.Default()
	r.GET("/users", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"ok": users,
		})
		spew.Dump(users)
	})
	r.Run()
}
