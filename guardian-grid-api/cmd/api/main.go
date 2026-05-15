package main

import (
	"guardian-grid-api/internal/modules/auth"
	"guardian-grid-api/internal/platform/database"
	"guardian-grid-api/internal/platform/security"
	"guardian-grid-api/internal/router"
	"fmt"
)

func main() {
	// Init Databases
	database.InitSQLite()
	database.InitMongoDB()

	// First run check
	if auth.IsFirstRun() {
		auth.RunInitialSetup()
	}

	r := router.SetupRouter()
	
	port := security.GetPort()
	fmt.Printf("🚀 GuardianGrid API starting on :%s\n", port)
	
	r.Run(":" + port)
}
