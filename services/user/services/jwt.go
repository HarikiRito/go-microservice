package services

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
)

type JwtService struct {
	Secret []byte
}

func (s *JwtService) GenerateToken(jwtMap jwt.MapClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtMap)
	tokenString, err := token.SignedString(s.Secret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (s *JwtService) VerifyToken(tokenString string) (map[string]interface{}, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return s.Secret, nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}
