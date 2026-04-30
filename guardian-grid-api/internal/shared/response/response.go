package response

import "github.com/gin-gonic/gin"

type APIResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

// Success response
func Success(c *gin.Context, data interface{}) {
	c.JSON(200, APIResponse{
		Status: "success",
		Data:   data,
	})
}

// Error response
func Error(c *gin.Context, message string) {
	c.JSON(400, APIResponse{
		Status:  "error",
		Message: message,
	})
}

func ErrorWithCode(c *gin.Context, code int, message string) {
	c.JSON(code, APIResponse{
		Status:  "error",
		Message: message,
	})
}
