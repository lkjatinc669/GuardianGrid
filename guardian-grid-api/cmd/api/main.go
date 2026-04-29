// cmd/api/main.go
package main

import (
	"guardian-grid-api/internal/db"
	"guardian-grid-api/internal/models"
	"guardian-grid-api/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Init databases
	db.InitSQLite()
	// db.InitMySQL()

	// Migrate tables
	db.SQLite.AutoMigrate(&models.User{})

	// Create server
	r := gin.Default()

	// Setup routes
	routes.SetupRoutes(r)

	// Run server
	r.Run(":64289")
}
