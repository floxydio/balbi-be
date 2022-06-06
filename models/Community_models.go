package models

import "time"

type Komunitas struct {
	Id             uint      `json:"id" gorm:"primaryKey"`
	Nama           string    `json:"nama" form:"nama"`
	Detail         string    `json:"detail" form:"detail"`
	Contact        string    `json:"contact" form:"contact"`
	PhotoKomunitas string    `json:"photo_komunitas" form:"photo_komunitas" gorm:"default:community_default.jpg"`
	CreatedBy      string    `json:"created_by" form:"created_by"`
	CreatedAt      time.Time `json:"created_at" form:"created_at"`
}

func (Komunitas) TableName() string {
	return "komunitas"
}
