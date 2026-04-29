package routes

import (
	"guardian-grid-api/internal/handlers"
	"guardian-grid-api/internal/handlers/dashboard"
	"guardian-grid-api/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {

	r.Any("/", handlers.HelloSimple)
	r.Any("/gg", handlers.Hello)

	// ========================================
	// API ROOT
	// ========================================

	api := r.Group("/dashboard-api")

	// ========================================
	// AUTH ROUTES (Public)
	// ========================================

	auth := api.Group("/auth")
	{
		auth.POST("/register", dashboard.Register)
		auth.POST("/login", dashboard.Login)
	}

	// ========================================
	// DASHBOARD ROUTES (Protected)
	// ========================================

	dash := api.Group("/dashboard")
	dash.Use(middleware.Protect())
	{
		dash.GET("/", dashboard.Dashboard)
	}

	// ========================================
	// AGENT ROUTES
	// ========================================
	// (add later)
}
