package models

type Task struct {
	BaseModel
	Name        string `gorm:"size:20" json:"name"`
	Theme       string `gorm:"size:200" json:"theme"`
	Description string `gorm:"size:2000" json:"description"`

	CreatorID uint `json:"creatorID"`
	Creator   User `gorm:"foreignKey:CreatorID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"creator"`

	ExecutorID uint `json:"executorID"`
	Executor   User `gorm:"foreignKey:ExecutorID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"executor"`

	PriorityID uint     `json:"priorityID"`
	Priority   Priority `gorm:"foreignKey:PriorityID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"priority"`

	StatusID uint   `json:"statusID"`
	Status   Status `gorm:"foreignKey:StatusID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"status"`
}
