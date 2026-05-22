package alerts

import (
	"github.com/gin-gonic/gin"
)

func RegisterDashboardRoutes(rg *gin.RouterGroup, h *Handler) {
	rg.GET("/alerts", h.GetDashboardAlerts)
}

func RegisterAgentRoutes(rg *gin.RouterGroup, h *Handler) {
	rg.GET("/alerts", h.GetAgentAlerts)
}
