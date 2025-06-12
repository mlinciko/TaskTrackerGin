package i_repositories

import (
	requests "gin/task_tracker/handlers/dtos/request"
	"gin/task_tracker/models"
)

type UserRepository interface {
	CreateUser(user *models.User) error
	GetUserByID(id uint) (*models.User, error)
	GetAllUsers(filter requests.GetAllUsersRequestDto) ([]*models.User, error)
	UpdateUser(user *models.User) (*models.User, error)
	DeleteUser(user *models.User) error
}
