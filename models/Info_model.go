package models

type InfoBlog struct {
	Id        uint   `json:"id" gorm:"primary_key"`
	Title     string `json:"title" form:"title"`
	Author    string `json:"author" form:"author"`
	Tag       string `json:"tag" form:"tag"`
	Deskripsi string `json:"deskripsi" form:"deskripsi"`
	Content   string `json:"content" form:"content"`
	Image     string `json:"image" form:"image"`
	Like      uint   `json:"like" form:"like"`
	CreatedAt string `json:"created_at" form:"created_at"`
}

func (InfoBlog) TableName() string {
	return "info_blog"
}
