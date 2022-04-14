package models

type ReviewParenting struct {
	Id          uint   `json:"id" gorm:"primaryKey"`
	ParentingID uint   `json:"parenting_id" form:"parenting_id"`
	UserID      uint   `json:"user_id" form:"user_id"`
	Rating      uint   `json:"rating" form:"rating"`
	Comment     string `json:"comment" form:"comment"`
}

func (ReviewParenting) TableName() string {
	return "review_parenting"
}
