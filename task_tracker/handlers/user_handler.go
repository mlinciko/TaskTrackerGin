package handlers

import (
	"fmt"
	requests "gin/task_tracker/handlers/dtos/request"
	user_errors "gin/task_tracker/handlers/errors"
	i_handlers "gin/task_tracker/handlers/interfaces"
	"gin/task_tracker/models"
	i_services "gin/task_tracker/services/interfaces"
	"gin/task_tracker/utils"
	"net/http"
	"strconv"

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
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, utils.MakeResp(nil, http.StatusBadRequest, err.Error()))
		return
	}

	user := models.User{FirstName: request.FirstName}

	if err := h.service.CreateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, utils.MakeResp(nil, http.StatusInternalServerError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, utils.MakeResp(user, http.StatusOK, ""))
}

func (h *userHandler) GetUserHandler(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.MakeResp(nil, http.StatusBadRequest, user_errors.InvalidUserId))
		return
	}

	user, err := h.service.GetUserByID(uint(id))
	if err != nil {
		errMessage := fmt.Sprintf(user_errors.UserNotFound, id)
		c.JSON(http.StatusNotFound, utils.MakeResp(nil, http.StatusNotFound, errMessage))
		return
	}

	c.JSON(http.StatusOK, utils.MakeResp(user, http.StatusOK, ""))
}

func (h *userHandler) PutUserHandler(c *gin.Context) {
	var request requests.PutUserRequestDto
	if err := c.ShouldBind(&request); err == nil {
		errMessage := fmt.Sprintf(user_errors.InvalidRequestBody, utils.GetFieldNames(requests.PutUserRequestDto{}))
		c.JSON(http.StatusBadRequest, utils.MakeResp(nil, http.StatusBadRequest, errMessage))
		return
	}

	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.MakeResp(nil, http.StatusBadRequest, user_errors.InvalidUserId))
		return
	}

	user, err := h.service.GetUserByID(uint(id))
	if err != nil {
		errMessage := fmt.Sprintf(user_errors.UserNotFound, id)
		c.JSON(http.StatusNotFound, utils.MakeResp(nil, http.StatusNotFound, errMessage))
		return
	}

	user.FirstName = request.FirstName

	user, err = h.service.UpdateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.MakeResp(nil, http.StatusInternalServerError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, utils.MakeResp(user, http.StatusOK, ""))
}

func (h *userHandler) PatchUserHandler(c *gin.Context) {
	var request requests.PatchUserRequestDto
	if err := c.ShouldBind(&request); err == nil {
		errMessage := fmt.Sprintf(user_errors.InvalidRequestBody, utils.GetFieldNames(requests.PatchUserRequestDto{}))
		c.JSON(http.StatusBadRequest, utils.MakeResp(nil, http.StatusBadRequest, errMessage))
		return
	}

	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.MakeResp(nil, http.StatusBadRequest, user_errors.InvalidUserId))
		return
	}

	user, err := h.service.GetUserByID(uint(id))
	if err != nil {
		errMessage := fmt.Sprintf(user_errors.UserNotFound, id)
		c.JSON(http.StatusNotFound, utils.MakeResp(nil, http.StatusNotFound, errMessage))
		return
	}

	if request.FirstName != "" {
		user.FirstName = request.FirstName
	}

	user, err = h.service.UpdateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.MakeResp(nil, http.StatusInternalServerError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, utils.MakeResp(user, http.StatusOK, ""))
}

func (h *userHandler) DeleteUserHandler(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.MakeResp(nil, http.StatusBadRequest, user_errors.InvalidUserId))
		return
	}

	user, err := h.service.GetUserByID(uint(id))
	if err != nil {
		errMessage := fmt.Sprintf(user_errors.UserNotFound, id)
		c.JSON(http.StatusNotFound, utils.MakeResp(nil, http.StatusNotFound, errMessage))
		return
	}

	err = h.service.DeleteUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.MakeResp(nil, http.StatusInternalServerError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, utils.MakeResp(user, http.StatusOK, ""))
}

func (h *userHandler) GetAllUsersHandler(c *gin.Context) {
	var request requests.GetAllUsersRequestDto
	if err := c.ShouldBind(&request); err == nil {
		errMessage := fmt.Sprintf(user_errors.InvalidRequestBody, utils.GetFieldNames(requests.GetAllUsersRequestDto{}))
		c.JSON(http.StatusBadRequest, utils.MakeResp(nil, http.StatusBadRequest, errMessage))
		return
	}

	users, err := h.service.GetAllUsers(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.MakeResp(nil, http.StatusInternalServerError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, utils.MakeResp(users, http.StatusOK, ""))
}
