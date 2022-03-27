package database

import (
	"balbibe/models"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	db, err := gorm.Open(mysql.Open("root:@tcp(localhost)/balbi?parseTime=true"), &gorm.Config{})
	if err != nil {
		panic("DB Not Connected")
	} else {
		fmt.Println("Connected")
	}

	DB = db

	// Migrate model to DB

	db.AutoMigrate(models.User{})
	db.AutoMigrate(models.Konsul{})
	db.AutoMigrate(models.Coupon{})
	db.AutoMigrate(models.Book{})
	db.AutoMigrate(models.InfoBlog{})
	db.AutoMigrate(models.Ads{})
	db.AutoMigrate(models.Course{})
	db.AutoMigrate(models.TransactionBook{})
}
