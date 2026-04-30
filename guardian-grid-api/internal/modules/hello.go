package handlers

import "github.com/gin-gonic/gin"

type HelloResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

func HelloSimple(c *gin.Context) {
	c.JSON(200, HelloResponse{
		Status:  true,
		Message: "GuardianGrid API is up",
	})
}

func Hello(c *gin.Context) {
	c.JSON(200, HelloResponse{
		Status:  true,
		Message: "Hello from GuardianGrid",
	})
}
