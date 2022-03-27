package models

type Ads struct {
	Id        uint   `json:"id"`
	Title     string `json:"title"`
	Image     string `json:"image"`
	Detail    string `json:"detail"`
	View      uint   `json:"view"`
	CreatedAt string `json:"created_at"`
}
