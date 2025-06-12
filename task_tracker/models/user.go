package models

type User struct {
	BaseModel `json:"baseModel"`
	FirstName string `gorm:"size:100" json:"firstName" binding:"required"`
}
