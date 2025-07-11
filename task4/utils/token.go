package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var JwtKey = "secret-key"

func GenerateToken(UserID int) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": UserID,
		"exp":     time.Now().Add(2 * time.Hour).Unix(),
	})
	tokenStr, _ := token.SignedString([]byte(JwtKey))
	return tokenStr
}
