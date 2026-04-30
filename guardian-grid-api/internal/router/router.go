package router

import (
	"guardian-grid-api/internal/modules/auth"
	"guardian-grid-api/internal/shared/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Auth routes (public)
	auth.RegisterRoutes(r)

	// Dashboard (JWT protected)
	dashboardGroup := r.Group("/dashboard")
	dashboardGroup.Use(middleware.JWTAuthMiddleware())
	{
		// dashboard.RegisterRoutes(dashboardGroup)
	}

	// Agent routes
	agentRepo := auth.NewAgentRepository()
	agentService := auth.NewAgentService(agentRepo)
	agentHandler := auth.NewAgentHandler(agentService)

	agentGroup := r.Group("/agent")
	{
		// public
		agentGroup.POST("/register", agentHandler.Register)

		// protected
		protected := agentGroup.Group("/")
		protected.Use(middleware.AgentAuthMiddleware())
		{
			protected.POST("/data", agentHandler.)
		}
	}

	return r
}
