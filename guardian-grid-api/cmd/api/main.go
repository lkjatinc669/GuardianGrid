package main

import (
	"guardian-grid-api/internal/modules/auth"
	"guardian-grid-api/internal/platform/database"
	"guardian-grid-api/internal/router"
)

func main() {
	// Init SQLite
	database.InitSQLite("static/guardian.db")

	// First run check
	if auth.IsFirstRun() {
		auth.RunInitialSetup()
	}

	r := router.SetupRouter()
	r.Run(":8080")

	// start your router here
}
