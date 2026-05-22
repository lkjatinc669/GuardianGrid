package alerts

import (
	"guardian-grid-api/internal/modules/websocket"
	"guardian-grid-api/internal/shared/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *Service
	hub     *websocket.Hub
}

func NewHandler(service *Service, hub *websocket.Hub) *Handler {
	return &Handler{service: service, hub: hub}
}

func (h *Handler) GetDashboardAlerts(c *gin.Context) {
	alerts, err := h.service.GetAllAlerts()
	if err != nil {
		response.Error(c, "Failed to fetch alerts")
		return
	}
	response.Success(c, alerts)
}

func (h *Handler) GetAgentAlerts(c *gin.Context) {
	agentID, _ := c.Get("agent_id")
	alerts, err := h.service.GetAgentAlerts(agentID.(string))
	if err != nil {
		response.Error(c, "Failed to fetch alerts")
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "data": alerts})
}
