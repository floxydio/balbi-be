package models

type KomunitasPost struct {
	Id          uint64 `json:"id" gorm:"primaryKey"`
	KomunitasID uint64 `json:"komunitas_id" form:"komunitas_id"`
	Title       string `json:"title" form:"title"`
	Content     string `json:"content" form:"content"`
	CreatedBy   uint64 `json:"created_by" form:"created_by"`
	CreatedAt   string `json:"created_at" form:"created_at"`
}

func (KomunitasPost) TableName() string {
	return "komunitas_post"
}
