package models

type Article struct {
	Id        uint   `json:"id" gorm:"primaryKey"`
	Title     string `json:"title" form:"title" gorm:"unique"`
	Content   string `json:"content" form:"content"`
	Author    string `json:"author" form:"author"`
	Rating    uint   `json:"rating" form:"rating" default:"0"`
	CreatedAt string `json:"created_at" form:"created_at"`
}

func (Article) TableName() string {
	return "article"
}
