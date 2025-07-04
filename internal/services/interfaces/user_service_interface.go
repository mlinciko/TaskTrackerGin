package i_services

import (
	requests "github.com/mlinciko/TaskTrackerGin/internal/handlers/dtos/request"
	"github.com/mlinciko/TaskTrackerGin/internal/models"
)

type UserService interface {
	CreateUser(user *models.User) error
	GetUserByID(id uint) (*models.User, error)
	GetAllUsers(filter requests.GetAllUsersRequestDto) ([]*models.User, error)
	UpdateUser(user *models.User) (*models.User, error)
	DeleteUser(user *models.User) error
	GetUserByEmail(email string) (*models.User, error)
}
