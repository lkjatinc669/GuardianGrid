package auth

import (
	"fmt"
	"guardian-grid-api/internal/modules/websocket"
	"guardian-grid-api/internal/shared/response"

	"github.com/gin-gonic/gin"
)

type AHandler struct {
	service *AService
	hub     *websocket.Hub
}

func NewAgentHandler(service *AService, hub *websocket.Hub) *AHandler {
	return &AHandler{service: service, hub: hub}
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

func (h *AHandler) PostData(c *gin.Context) {
	agentID, _ := c.Get("agent_id")
	
	var payload map[string]interface{}
	if err := c.ShouldBindJSON(&payload); err != nil {
		response.Error(c, "Invalid data format")
		return
	}

	err := h.service.ProcessTelemetry(fmt.Sprintf("%v", agentID), payload)
	if err != nil {
		response.Error(c, "Failed to save telemetry")
		return
	}

	// 📡 Stream to dashboard
	if h.hub != nil {
		h.hub.BroadcastTelemetry(fmt.Sprintf("%v", agentID), payload)
	}

	fmt.Printf("📥 Telemetry saved for agent: %v\n", agentID)
	response.Success(c, gin.H{"status": "received"})
}
