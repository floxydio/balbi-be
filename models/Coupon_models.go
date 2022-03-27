package models

type Coupon struct {
	Id        uint   `json:"id" gorm:"primaryKey" form:"id"`
	Token     string `json:"token" form:"token"`
	ExpiredAt string `json:"expiredAt" form:"expiredAt"`
}
