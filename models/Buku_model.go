package models

type Book struct {
	Id         uint    `json:"id" gorm:"primary_key"`
	Title      string  `json:"title" form:"title"`
	Rating     float64 `json:"rating" form:"rating"`
	Page       uint    `json:"page" form:"page"`
	Price      uint    `json:"price" form:"price"`
	File       string  `json:"file" form:"file"`
	AlreadyBuy uint    `json:"alreadybuy" form:"alreadybuy" default:"0"`
	Author     string  `json:"author" form:"author"`
}

func (Book) TableName() string {
	return "book"
}
