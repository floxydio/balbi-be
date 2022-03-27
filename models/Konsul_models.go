package models

type Konsul struct {
	Id        uint   `json:"id" gorm:"primaryKey" form:"id"`
	Nama      string `json:"nama" form:"nama" validate:"required"`
	NamaUser  string `json:"nama_user" form:"nama_user" validate:"required"`
	Status    uint   `json:"status" form:"status" default:"0"`
	Lokasi    string `json:"lokasi" form:"lokasi" validate:"required"`
	Keluhan   string `json:"keluhan" form:"keluhan" validate:"required"`
	Message   string `json:"message" validate:"required,min=10" form:"message" default:"-"`
	BookingAt string `json:"booking_at" form:"booking_at"`
	CreatedAt string `json:"created_at" form:"created_at"`
}

func (Konsul) TableName() string {
	return "konsul"
}
