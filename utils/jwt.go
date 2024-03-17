package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

func SignJwt(email string) (string, error) {
	expirationTime := time.Now().Add(30 * 1440 * time.Minute)

	claims := &Claims{
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_TOKEN")))

	return tokenString, err
}

func ParseJwt(jwtToken string) (*Claims, *jwt.Token, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(jwtToken, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_TOKEN")), nil
	})

	return claims, token, err
}
