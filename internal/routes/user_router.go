package routes

import (
	"github.com/mlinciko/TaskTrackerGin/internal/database"
	"github.com/mlinciko/TaskTrackerGin/internal/handlers"
	"github.com/mlinciko/TaskTrackerGin/internal/middleware"
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

	userGroupProtected := userGroup.Use(middleware.VerifyAccess())
	userGroupProtected.GET("/:id", handler.GetUserHandler)
	userGroupProtected.PUT("/:id", handler.PutUserHandler)
	userGroupProtected.PATCH("/:id", handler.PatchUserHandler)
	userGroupProtected.DELETE("/:id", handler.DeleteUserHandler)
	userGroupProtected.GET("/all", handler.GetAllUsersHandler)
}
