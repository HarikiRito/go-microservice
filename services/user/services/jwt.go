package services

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

type JwtService struct {
}

var secret = "sayang"

func (s *JwtService) GenerateToken(jwtMap jwt.MapClaims) (string, error) {
	jwtMap["createdAt"] = time.Now().Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtMap)
	secretBytes := []byte(secret)
	tokenString, err := token.SignedString(secretBytes)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
