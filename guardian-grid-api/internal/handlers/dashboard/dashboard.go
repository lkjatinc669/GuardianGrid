// internal/handlers/dashboard.go
package dashboard

import "github.com/gin-gonic/gin"

func Dashboard(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "dashboard working",
	})
}
