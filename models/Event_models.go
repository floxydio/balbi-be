package models

import "time"

type Event struct {
	Id        uint      `json:"id" gorm:"primaryKey"`
	Image     string    `json:"image" form:"image"`
	Detail    string    `json:"detail" form:"detail"`
	Winner    string    `json:"winner" form:"winner"`
	LinkEvent string    `json:"link_event" form:"link_event"`
	CreatedAt time.Time `json:"created_at"`
}

func (Event) TableName() string {
	return "event"
}
