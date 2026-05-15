package dashboard

import "github.com/gin-gonic/gin"

func RegisterRoutes(rg *gin.RouterGroup, h *Handler) {
	rg.GET("/overview", h.GetOverview)
	rg.GET("/agent/:id", h.GetAgent)
}
