package router

import (
	"guardian-grid-api/internal/modules/auth"
	"guardian-grid-api/internal/modules/dashboard"
	"guardian-grid-api/internal/modules/websocket"
	"guardian-grid-api/internal/shared/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Apply CORS
	r.Use(middleware.CORSMiddleware())

	// WebSocket Hub
	hub := websocket.NewHub()
	go hub.Run()

	// Auth routes (public)
	auth.RegisterRoutes(r, hub)

	// Stream (Dashboard WebSocket)
	r.GET("/ws", func(c *gin.Context) {
		hub.ServeHTTP(c.Writer, c.Request)
	})

	// Dashboard (JWT protected)
	dashRepo := dashboard.NewRepository()
	dashService := dashboard.NewService(dashRepo)
	dashHandler := dashboard.NewHandler(dashService)

	dashboardGroup := r.Group("/dashboard")
	dashboardGroup.Use(middleware.JWTAuthMiddleware())
	{
		dashboard.RegisterRoutes(dashboardGroup, dashHandler)
	}

	// Agent routes
	agentRepo := auth.NewAgentRepository()
	agentService := auth.NewAgentService(agentRepo)
	agentHandler := auth.NewAgentHandler(agentService, hub)

	agentGroup := r.Group("/agent")
	{
		// public
		agentGroup.POST("/register", agentHandler.Register)

		// protected
		protected := agentGroup.Group("/")
		protected.Use(middleware.AgentAuthMiddleware())
		{
			protected.POST("/data", agentHandler.PostData)
		}
	}

	return r
}
