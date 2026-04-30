package middleware

import (
	"net/http"
	"strings"

	"guardian-grid-api/internal/modules/auth"
	"guardian-grid-api/internal/platform/security"
	"guardian-grid-api/internal/shared/response"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		// Expect: Bearer <token>
		if authHeader == "" {
			response.ErrorWithCode(c, http.StatusUnauthorized, "Missing Authorization header")
			c.Abort()
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			response.ErrorWithCode(c, http.StatusUnauthorized, "Invalid Authorization format")
			c.Abort()
			return
		}

		tokenStr := parts[1]

		// Parse token
		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			return security.GetJWTSecret(), nil
		})

		if err != nil || !token.Valid {
			response.ErrorWithCode(c, http.StatusUnauthorized, "Invalid or expired token")
			c.Abort()
			return
		}

		// Extract claims
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			response.ErrorWithCode(c, http.StatusUnauthorized, "Invalid token claims")
			c.Abort()
			return
		}

		// Attach user to context
		user, _ := claims["user"].(string)
		c.Set("user", user)

		c.Next()
	}
}

func AgentAuthMiddleware() gin.HandlerFunc {
	repo := auth.NewAgentRepository()

	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			response.ErrorWithCode(c, http.StatusUnauthorized, "Missing Authorization header")
			c.Abort()
			return
		}

		// Expect: Bearer <token>
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			response.ErrorWithCode(c, http.StatusUnauthorized, "Invalid Authorization format")
			c.Abort()
			return
		}

		token := parts[1]

		agentID, err := repo.GetAgentByToken(token)
		if err != nil {
			response.ErrorWithCode(c, http.StatusUnauthorized, "Invalid agent token")
			c.Abort()
			return
		}

		// Attach agent_id to request context
		c.Set("agent_id", agentID)

		c.Next()
	}
}
