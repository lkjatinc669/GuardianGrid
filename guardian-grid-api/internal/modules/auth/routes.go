package auth

import (
	"guardian-grid-api/internal/modules/websocket"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, hub *websocket.Hub) {
	// user auth
	userRepo := NewDashRepository()
	userService := NewDashService(userRepo)
	userHandler := NewDashHandler(userService)

	// agent auth
	agentRepo := NewAgentRepository()
	agentService := NewAgentService(agentRepo)
	agentHandler := NewAgentHandler(agentService, hub)

	group := r.Group("/auth")
	{
		group.POST("/login", userHandler.Login)

		// 👇 agent registration
		group.POST("/agent/register", agentHandler.Register)
	}
}
