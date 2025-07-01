package factories

import (
	"time"

	"github.com/mlinciko/TaskTrackerGin/internal/models"

	"github.com/brianvoe/gofakeit/v6"
)

func CreateUser() *models.User {
	return &models.User{
		BaseModel: models.BaseModel{ID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		FirstName: gofakeit.FirstName(),
		Email:     gofakeit.Email(),
	}
}
