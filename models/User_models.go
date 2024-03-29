package models

type User struct {
	Id           uint   `json:"id" gorm:"primaryKey" form:"id"`
	Nama         string `json:"nama" form:"nama" validate:"required"`
	Asal         string `json:"asal" form:"asal" default:"-"`
	Email        string `json:"email" form:"email" validate:"required,email" default:"-"`
	Username     string `json:"username" form:"username"`
	Password     string `json:"password" form:"password" default:"-"`
	EmailValid   string `json:"email_verif" form:"email_verif" default:"0"`
	Phone        uint   `json:"phone" form:"phone" gorm:"column:phone"`
	Photo        string `json:"photo" form:"photo" default:"default.jpg"`
	Role         uint   `json:"role" form:"role" default:"0"`
	Premium      uint   `json:"premium" form:"premium" default:"0"`
	Point        uint   `json:"point" form:"point" default:"0"`
	IdPertanyaan uint   `json:"id_pertanyaan" form:"idpertanyaan" gorm:"column:idpertanyaan"`
	Jawaban      string `json:"jawaban" form:"jawaban" default:"-" gorm:"column:jawaban"`
	CreatedAt    string `json:"created_at" form:"created_at"`
	UpdatedAt    string `json:"updated_at" form:"updated_at"`
}

func (User) TableName() string {
	return "user"
}
