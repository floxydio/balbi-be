package models

type Notification struct {
	Id      uint   `json:"id" gorm:"primaryKey"`
	Message string `json:"message" form:"message"`
}

func (Notification) TableName() string {
	return "notification"
}
