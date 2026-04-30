package auth

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.Engine) {
	// user auth
	userRepo := DashRepository()
	userService := DashService(userRepo)
	userHandler := DashHandler(userService)

	// agent auth
	agentRepo := AgentRepository()
	agentService := AgentService(agentRepo)
	agentHandler := AgentHandler(agentService)

	group := r.Group("/auth")
	{
		group.POST("/login", userHandler.Login)

		// 👇 agent registration
		group.POST("/agent/register", agentHandler.Register)
	}
}
