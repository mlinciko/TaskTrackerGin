package routes

import (
	"gin/task_tracker/handlers"

	"github.com/gin-gonic/gin"
)

func SetupUserGroup(router *gin.Engine) {
	userGroup := router.Group("/users")
	userGroup.POST("", handlers.CreateUserHandler)
	userGroup.GET("/:id", handlers.GetUserHandler)
	userGroup.PUT("/:id", handlers.PutUserHandler)
	userGroup.PATCH("/:id", handlers.PatchUserHandler)
	userGroup.DELETE("/:id", handlers.DeleteUserHandler)
	userGroup.GET("/employees", handlers.GetAllUsersHandler)
}
