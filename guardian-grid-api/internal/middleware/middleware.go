// internal/middleware/auth.go
package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var secret = []byte("change-this-later")

func Protect() gin.HandlerFunc {
	return func(c *gin.Context) {

		header := c.GetHeader("Authorization")

		if header == "" {
			c.AbortWithStatusJSON(401, gin.H{"error": "no token"})
			return
		}

		tokenString := strings.Split(header, " ")[1]

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return secret, nil
		})

		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(401, gin.H{"error": "invalid token"})
			return
		}

		c.Next()
	}
}
