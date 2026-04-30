package auth

import (
	"guardian-grid-api/internal/shared/response"

	"github.com/gin-gonic/gin"
)

type AHandler struct {
	service *AService
}

func AgentHandler(service *AService) *AHandler {
	return &AHandler{service: service}
}

func (h *AHandler) Register(c *gin.Context) {
	var req struct {
		Hostname string `json:"hostname"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, "Invalid request")
		return
	}

	id, token, err := h.service.RegisterAgent(req.Hostname)
	if err != nil {
		response.Error(c, "Registration failed")
		return
	}

	response.Success(c, gin.H{
		"agent_id": id,
		"token":    token,
	})
}
