package auth

import (
	"guardian-grid-api/internal/platform/security"
	"guardian-grid-api/internal/shared/response"

	"github.com/gin-gonic/gin"
)

type DHandler struct {
	service *DService
}

func DashHandler(service *DService) *DHandler {
	return &DHandler{service: service}
}

func (h *DHandler) Login(c *gin.Context) {
	var req LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, "Invalid request")
		return
	}

	user, err := h.service.Login(req.Username, req.Code)
	if err != nil {
		response.Error(c, err.Error())
		return
	}

	token, err := security.GenerateJWT(user.Username)
	if err != nil {
		response.Error(c, "Failed to generate token")
		return
	}

	response.Success(c, gin.H{
		"token": token,
	})
}
