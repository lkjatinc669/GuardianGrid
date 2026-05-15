package dashboard

import (
	"guardian-grid-api/internal/shared/response"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) GetOverview(c *gin.Context) {
	data, err := h.service.GetGlobalStatus()
	if err != nil {
		response.Error(c, "Failed to fetch overview")
		return
	}
	response.Success(c, data)
}

func (h *Handler) GetAgent(c *gin.Context) {
	id := c.Param("id")
	data, err := h.service.GetAgentDetails(id)
	if err != nil {
		response.Error(c, "Agent not found")
		return
	}
	response.Success(c, data)
}
