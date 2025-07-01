package handlers

import (
	"net/http"
	"strconv"

	xerrors "github.com/mlinciko/TaskTrackerGin/internal/errors"
	requests "github.com/mlinciko/TaskTrackerGin/internal/handlers/dtos/request"
	i_handlers "github.com/mlinciko/TaskTrackerGin/internal/handlers/interfaces"
	"github.com/mlinciko/TaskTrackerGin/internal/models"
	i_services "github.com/mlinciko/TaskTrackerGin/internal/services/interfaces"
	"github.com/mlinciko/TaskTrackerGin/internal/utils"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	service i_services.UserService
}

func NewUserHandler(service i_services.UserService) i_handlers.UserHandler {
	return &userHandler{service}
}

func (h *userHandler) CreateUserHandler(c *gin.Context) {
	var request requests.CreateUserRequestDto
	if err := c.ShouldBind(&request); err != nil {
		c.Error(err).SetMeta(http.StatusBadRequest)
		return
	}

	user := models.User{FirstName: request.FirstName, Email: request.Email}

	if err := h.service.CreateUser(&user); err != nil {
		c.Error(err).SetMeta(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, utils.MakeResp(user, http.StatusOK, ""))
}

func (h *userHandler) GetUserHandler(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.Error(xerrors.ErrInvalidUserId).SetMeta(http.StatusBadRequest)
		return
	}

	user, err := h.service.GetUserByID(uint(id))
	if err != nil {
		c.Error(xerrors.ErrUserNotFound).SetMeta(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, utils.MakeResp(user, http.StatusOK, ""))
}

func (h *userHandler) PutUserHandler(c *gin.Context) {
	var request requests.PutUserRequestDto
	if err := c.ShouldBind(&request); err == nil {
		c.Error(err).SetMeta(http.StatusBadRequest)
		return
	}

	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.Error(xerrors.ErrInvalidUserId).SetMeta(http.StatusBadRequest)
		return
	}

	user, err := h.service.GetUserByID(uint(id))
	if err != nil {
		c.Error(xerrors.ErrUserNotFound).SetMeta(http.StatusNotFound)
		return
	}

	user.FirstName = request.FirstName
	user.Email = request.Email

	user, err = h.service.UpdateUser(user)
	if err != nil {
		c.Error(err).SetMeta(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, utils.MakeResp(user, http.StatusOK, ""))
}

func (h *userHandler) PatchUserHandler(c *gin.Context) {
	var request requests.PatchUserRequestDto
	if err := c.ShouldBind(&request); err == nil {
		c.Error(err).SetMeta(http.StatusBadRequest)
		return
	}

	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.Error(xerrors.ErrInvalidUserId).SetMeta(http.StatusBadRequest)
		return
	}

	user, err := h.service.GetUserByID(uint(id))
	if err != nil {
		c.Error(xerrors.ErrUserNotFound).SetMeta(http.StatusNotFound)
		return
	}

	if request.FirstName != "" {
		user.FirstName = request.FirstName
	}

	if request.Email != "" {
		user.Email = request.Email
	}

	user, err = h.service.UpdateUser(user)
	if err != nil {
		c.Error(err).SetMeta(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, utils.MakeResp(user, http.StatusOK, ""))
}

func (h *userHandler) DeleteUserHandler(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.Error(xerrors.ErrInvalidUserId).SetMeta(http.StatusBadRequest)
		return
	}

	user, err := h.service.GetUserByID(uint(id))
	if err != nil {
		c.Error(xerrors.ErrUserNotFound).SetMeta(http.StatusNotFound)
		return
	}

	err = h.service.DeleteUser(user)
	if err != nil {
		c.Error(err).SetMeta(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, utils.MakeResp(user, http.StatusOK, ""))
}

func (h *userHandler) GetAllUsersHandler(c *gin.Context) {
	var request requests.GetAllUsersRequestDto
	if err := c.ShouldBind(&request); err != nil {
		c.Error(err).SetMeta(http.StatusBadRequest)
		return
	}

	users, err := h.service.GetAllUsers(request)
	if err != nil {
		c.Error(err).SetMeta(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, utils.MakeResp(users, http.StatusOK, ""))
}
