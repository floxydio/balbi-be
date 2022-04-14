package models

type AnakPertanyaan struct {
	Id                uint   `json:"id" gorm:"primaryKey"`
	UserID            uint   `json:"user_id" form:"user_id"`
	RatingAnakPertama uint   `json:"rating_anak_pertama" form:"rating_anak_pertama"`
	NamaAnakPertama   string `json:"nama_anak_pertama" form:"nama_anak_pertama"`
	RatingAnakKedua   uint   `json:"rating_anak_kedua" form:"rating_anak_kedua"`
	NamaAnakKedua     string `json:"nama_anak_kedua" form:"nama_anak_kedua"`
	RatingAnakKetiga  uint   `json:"rating_anak_ketiga" form:"rating_anak_ketiga"`
	NamaAnakKetiga    string `json:"nama_anak_ketiga" form:"nama_anak_ketiga"`
	UpdateAt          string `json:"update_at" form:"update_at"`
}

func (AnakPertanyaan) TableName() string {
	return "anak_pertanyaan"
}
