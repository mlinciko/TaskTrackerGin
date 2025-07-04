package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	requests "github.com/mlinciko/TaskTrackerGin/internal/handlers/dtos/request"
	i_handlers "github.com/mlinciko/TaskTrackerGin/internal/handlers/interfaces"
	i_services "github.com/mlinciko/TaskTrackerGin/internal/services/interfaces"
	"github.com/mlinciko/TaskTrackerGin/internal/utils"
)

type authHandler struct {
	service i_services.AuthService
}

func NewAuthHandler(service i_services.AuthService) i_handlers.AuthHandler {
	return &authHandler{service}
}

func (h *authHandler) Login(c *gin.Context) {
	var request requests.LoginRequestDto
	if err := c.ShouldBind(&request); err != nil {
		c.Error(err)
		return
	}

	token, err := h.service.Login(request)
	if err != nil {
		c.Error(err).SetMeta(http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusOK, utils.MakeResp(token, http.StatusOK, ""))
}
