package models

type Priority struct {
	ID   uint   `gorm:"primarykey" json:"id"`
	Name string `gorm:"size:20" json:"name"`
}
