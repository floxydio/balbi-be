package models

type KomunitasUser struct {
	Id          uint64 `json:"id" gorm:"primaryKey"`
	KomunitasID uint64 `json:"komunitas_id" form:"komunitas_id"`
	UserID      uint64 `json:"user_id" form:"user_id"`
}

func (KomunitasUser) TableName() string {
	return "komunitas_user"
}
