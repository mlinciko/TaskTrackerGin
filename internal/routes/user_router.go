package routes

import (
	"github.com/mlinciko/TaskTrackerGin/internal/database"
	"github.com/mlinciko/TaskTrackerGin/internal/handlers"
	"github.com/mlinciko/TaskTrackerGin/internal/repositories"
	"github.com/mlinciko/TaskTrackerGin/internal/services"

	"github.com/gin-gonic/gin"

	i_handlers "github.com/mlinciko/TaskTrackerGin/internal/handlers/interfaces"
)

func SetupUserGroup(router *gin.Engine) {
	repo := repositories.NewUserRepository(database.DB)
	service := services.NewUserService(repo)
	handler := handlers.NewUserHandler(service)

	SetUsersRoutes(router, handler)
}

func SetUsersRoutes(router *gin.Engine, handler i_handlers.UserHandler) {
	userGroup := router.Group("/users")
	userGroup.POST("", handler.CreateUserHandler)
	userGroup.GET("/:id", handler.GetUserHandler)
	userGroup.PUT("/:id", handler.PutUserHandler)
	userGroup.PATCH("/:id", handler.PatchUserHandler)
	userGroup.DELETE("/:id", handler.DeleteUserHandler)
	userGroup.GET("/all", handler.GetAllUsersHandler)
}
