package services

import (
	requests "github.com/mlinciko/TaskTrackerGin/internal/handlers/dtos/request"
	"github.com/mlinciko/TaskTrackerGin/internal/models"
	i_repositories "github.com/mlinciko/TaskTrackerGin/internal/repositories/interfaces"
	i_services "github.com/mlinciko/TaskTrackerGin/internal/services/interfaces"
	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	repo i_repositories.UserRepository
}

func NewUserService(repo i_repositories.UserRepository) i_services.UserService {
	return &userService{repo}
}

func (s *userService) CreateUser(user *models.User) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(bytes)

	if err != nil {
		return err
	}

	return s.repo.CreateUser(user)
}

func (s *userService) GetUserByID(id uint) (*models.User, error) {
	return s.repo.GetUserByID(id)
}

func (s *userService) GetAllUsers(filter requests.GetAllUsersRequestDto) ([]*models.User, error) {
	return s.repo.GetAllUsers(filter)
}

func (s *userService) UpdateUser(user *models.User) (*models.User, error) {
	return s.repo.UpdateUser(user)
}

func (s *userService) DeleteUser(user *models.User) error {
	return s.repo.DeleteUser(user)
}

func (s *userService) GetUserByEmail(email string) (*models.User, error) {
	return s.repo.GetUserByEmail(email)
}
