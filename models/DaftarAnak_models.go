package models

type DaftarAnak struct {
	Id       uint   `json:"id" gorm:"primaryKey"`
	UserId   uint   `json:"userid" form:"userid" gorm:"column:userid"`
	Gender   uint   `json:"gender" form:"gender" gorm:"column:gender"`
	NamaAnak string `json:"namaanak" form:"namaanak" gorm:"column:namaanak"`
	NamaPanggilan string `json:"namapanggilan" form:"namapanggilan"`
}
