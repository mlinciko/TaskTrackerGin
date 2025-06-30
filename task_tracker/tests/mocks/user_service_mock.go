package mocks

import (
	"gin/task_tracker/models"

	requests "gin/task_tracker/handlers/dtos/request"

	"github.com/stretchr/testify/mock"
)

type MockUserService struct {
	mock.Mock
}

func (m *MockUserService) CreateUser(user *models.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserService) GetUserByID(id uint) (*models.User, error) {
	args := m.Called(id)

	user, ok := args.Get(0).(*models.User)
	if !ok && args.Get(0) != nil {
		panic("expected *models.User as first return value")
	}

	return user, args.Error(1)
}

func (m *MockUserService) GetAllUsers(filter requests.GetAllUsersRequestDto) ([]*models.User, error) {
	args := m.Called(filter)

	user, ok := args.Get(0).([]*models.User)
	if !ok && args.Get(0) != nil {
		panic("expected []*models.User as first return value")
	}

	return user, args.Error(1)
}

func (m *MockUserService) UpdateUser(user *models.User) (*models.User, error) {
	args := m.Called(user)

	user, ok := args.Get(0).(*models.User)
	if !ok && args.Get(0) != nil {
		panic("expected *models.User as first return value")
	}

	return user, args.Error(1)
}

func (m *MockUserService) DeleteUser(user *models.User) error {
	args := m.Called(user)
	return args.Error(0)
}
