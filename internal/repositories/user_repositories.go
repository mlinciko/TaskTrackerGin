package repositories

import (
	xerrors "github.com/mlinciko/TaskTrackerGin/internal/errors"
	requests "github.com/mlinciko/TaskTrackerGin/internal/handlers/dtos/request"
	"github.com/mlinciko/TaskTrackerGin/internal/models"
	i_repositories "github.com/mlinciko/TaskTrackerGin/internal/repositories/interfaces"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) i_repositories.UserRepository {
	return &userRepository{db}
}

func (r *userRepository) CreateUser(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) GetUserByID(id uint) (*models.User, error) {
	var user models.User
	err := r.db.First(&user, id).Error
	return &user, err
}

func (r *userRepository) GetAllUsers(filter requests.GetAllUsersRequestDto) ([]*models.User, error) {
	var users []*models.User
	err := r.db.Where(models.User{FirstName: filter.FirstName}).Find(&users).Error
	return users, err
}

func (r *userRepository) UpdateUser(user *models.User) (*models.User, error) {
	tx := r.db.Model(&user).Updates(models.User{FirstName: user.FirstName, Email: user.Email})

	if tx.Error != nil {
		return nil, tx.Error
	}

	if tx.RowsAffected == 0 {
		return nil, xerrors.ErrZeroRowsAffected
	}

	user, err := r.GetUserByID(user.ID)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *userRepository) DeleteUser(user *models.User) error {
	return r.db.Delete(&user).Error
}

func (r *userRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.db.Where(models.User{Email: email}).First(&user).Error
	return &user, err
}
