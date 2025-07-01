package i_handlers

import "github.com/gin-gonic/gin"

type UserHandler interface {
	CreateUserHandler(c *gin.Context)
	GetUserHandler(c *gin.Context)
	PutUserHandler(c *gin.Context)
	PatchUserHandler(c *gin.Context)
	DeleteUserHandler(c *gin.Context)
	GetAllUsersHandler(c *gin.Context)
}
