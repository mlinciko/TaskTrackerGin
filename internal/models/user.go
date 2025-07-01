package models

type User struct {
	BaseModel
	FirstName string `gorm:"size:100;not null" json:"firstName" binding:"required"`
	Email     string `gorm:"size:100;not null;unique" json:"email" binding:"required,email"`
}
