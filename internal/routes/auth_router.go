package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mlinciko/TaskTrackerGin/internal/database"
	"github.com/mlinciko/TaskTrackerGin/internal/handlers"
	i_handlers "github.com/mlinciko/TaskTrackerGin/internal/handlers/interfaces"
	"github.com/mlinciko/TaskTrackerGin/internal/repositories"
	"github.com/mlinciko/TaskTrackerGin/internal/services"
)

func SetupAuthGroup(router *gin.Engine) {
	repo := repositories.NewUserRepository(database.DB)
	userService := services.NewUserService(repo)
	service := services.NewAuthService(repo, userService)
	handler := handlers.NewAuthHandler(service)

	SetAuthRoutes(router, handler)
}

func SetAuthRoutes(router *gin.Engine, handler i_handlers.AuthHandler) {
	userGroup := router.Group("/auth")
	userGroup.POST("", handler.Login)
}
