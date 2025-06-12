package repositories

import (
	"fmt"
	"gin/task_tracker/database"
	requests "gin/task_tracker/handlers/dtos/request"
	"gin/task_tracker/handlers/errors"
	"gin/task_tracker/models"
)

func CreateUser(user *models.User) error {
	return database.DB.Create(user).Error
}

func GetUserByID(id uint) (*models.User, error) {
	var user models.User
	err := database.DB.First(&user, id).Error
	return &user, err
}

func GetAllUsers(filter requests.GetAllUsersRequestDto) ([]*models.User, error) {
	var users []*models.User
	err := database.DB.Where(models.User{FirstName: filter.FirstName}).Find(&users).Error
	return users, err
}

func UpdateUser(user *models.User) (*models.User, error) {
	tx := database.DB.Model(&user).Updates(models.User{FirstName: user.FirstName})

	if tx.Error != nil {
		return nil, tx.Error
	}

	if tx.RowsAffected == 0 {
		return nil, fmt.Errorf(errors.ZeroRowsAffected, "UpdateUser")
	}

	user, err := GetUserByID(user.ID)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func DeleteUser(user *models.User) error {
	return database.DB.Delete(&user).Error
}
