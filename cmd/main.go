package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mlinciko/TaskTrackerGin/config"
	"github.com/mlinciko/TaskTrackerGin/internal/database"
	"github.com/mlinciko/TaskTrackerGin/internal/routes"
	jwt_utils "github.com/mlinciko/TaskTrackerGin/internal/utils/jwt"
	"github.com/mlinciko/TaskTrackerGin/internal/validators"
)

func main() {
	// Generate JWT secret key
	jwt_utils.GenerateAndSaveJWTSecretKey()

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
