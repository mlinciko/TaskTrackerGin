package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mlinciko/TaskTrackerGin/config"
	"github.com/mlinciko/TaskTrackerGin/internal/database"
	"github.com/mlinciko/TaskTrackerGin/internal/routes"
	"github.com/mlinciko/TaskTrackerGin/internal/validators"
)

var db = make(map[string]string)

func main() {
	// Load config
	config.LoadConfig()

	// Connection to DB
	dsn := config.AppConfig.GetPostgresDSN()
	database.Connect(dsn)
	database.Migrate()

	// Register validators
	validators.RegisterValidators()

	// Setup routes
	r := gin.Default()
	routes.SetupRoutes(r)

	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
