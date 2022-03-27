package models

import "time"

type Course struct {
	Id         uint   `json:"id" form:"id" gorm:"primaryKey"`
	Thumbnail  string `json:"thumbnail" form:"thumbnail"`
	Title      string `json:"title" form:"title" validate:"required"`
	View       uint64 `json:"view" column:"view" default:"0"`
	Url        string `json:"url[0]" form:"url"`
	Url2       string `json:"url[1]" form:"url[1]" column:"url_1"`
	Url3       string `json:"url[2]" form:"url[2]" column:"url_2"`
	Url4       string `json:"url[3]" form:"url[3]" column:"url_3"`
	Url5       string `json:"url[4]" form:"url[4]" column:"url_4"`
	Url6       string `json:"url[5]" form:"url[5]" column:"url_5"`
	Url7       string `json:"url[6]" form:"url[6]" column:"url_6"`
	Url8       string `json:"url[7]" form:"url[7]" column:"url_7"`
	Url9       string `json:"url[8]" form:"url[8]" column:"url_8"`
	Like       uint64 `json:"like" form:"like" default:"0"`
	Uploadedby string `json:"uploadedby" form:"uploadedby"`
	JobsTitle  string `json:"jobstitle" form:"jobstitle"`
	Status     string `json:"status" form:"status"`
	Deskripsi  string `json:"deskripsi" form:"deskripsi"`
	Premium    uint   `json:"premium" form:"premium" default:"0"`
}

type HistoryTransactionCourse struct {
	Id        uint      `json:"id"`
	Nama      string    `json:"nama" form:"nama" column:"nama"`
	BookTitle string    `json:"booktitle" form:"booktitle" column:"book_title"`
	BuyAt     time.Time `json:"buyat" form:"buyat" column:"buy_at"`
	File      string    `json:"file" form:"file" column:"file"`
}

func (Course) TableName() string {
	return "course"
}
