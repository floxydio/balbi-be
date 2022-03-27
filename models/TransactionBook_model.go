package models

import "time"

type TransactionBook struct {
	Id        uint      `json:"id" gorm:"primaryKey"`
	IdUser    uint      `json:"iduser" form:"iduser"`
	BookTitle string    `json:"booktitle" form:"booktitle"`
	IdBook    uint      `json:"idbook" form:"idbook"`
	BuyAt     time.Time `json:"buy_at"`
}

func (TransactionBook) TableName() string {
	return "transaction_book"
}
