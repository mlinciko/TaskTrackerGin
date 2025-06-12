package services

import (
	requests "gin/task_tracker/handlers/dtos/request"
	"gin/task_tracker/models"
	i_repositories "gin/task_tracker/repositories/interfaces"
	i_services "gin/task_tracker/services/interfaces"
)

type userService struct {
	repo i_repositories.UserRepository
}

func NewUserService(repo i_repositories.UserRepository) i_services.UserService {
	return &userService{repo}
}

func (s *userService) CreateUser(user *models.User) error {
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
