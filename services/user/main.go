package main

import (
	"user/database"
	"user/rpc"
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
	//user := models.UserEntity{Name: "main", Age: 12, Password: "zzz"}
	//r := repositories.UserRepository{}
	//r.Create(user)

	rpc.RunRPCServer()
}
