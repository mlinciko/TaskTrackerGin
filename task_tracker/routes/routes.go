package routes

import (
	"gin/task_tracker/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.Use(middleware.ErrorHandler())
	SetupUserGroup(router)
}
