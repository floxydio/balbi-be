package models

type VideoPlace struct {
	Id        uint   `json:"id" gorm:"primaryKey"`
	Title     string `json:"title" form:"title" gorm:"title"`
	Subtitle  string `json:"subtitle" form:"subtitle"`
	Category  string `json:"category" form:"category"`
	Deskripsi string `json:"deskripsi" form:"deskripsi"`
	Detail    string `json:"detail" form:"detail"`
	Time      int    `json:"time" form:"time"`
	Rating    int    `json:"rating" form:"rating"`
	Kota      string `json:"kota" form:"kota"`
	Harga     uint   `json:"harga" form:"harga"`
	Url       string `json:"url" form:"url"`
}

func (VideoPlace) TableName() string {
	return "video_place"
}
