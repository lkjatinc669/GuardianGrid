// internal/auth/jwt.go
package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secret = []byte("supersecret")

func GenerateToken(userID uint) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	t, _ := token.SignedString(secret)
	return t
}
