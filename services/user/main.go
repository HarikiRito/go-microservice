package main

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"user/database"
	"user/repositories"
)

func main() {
	database.MigrateTable()

	//jwtService := services.JwtService{
	//	Secret: []byte("sayang"),
	//}
	// Sign and get the complete encoded token as a string using the secret
	//tokenString, _ := jwtService.GenerateToken(jwt.MapClaims{
	//	"userId": 648,
	//})

	hash, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)

	fmt.Println(string(hash))
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
