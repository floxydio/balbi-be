package database

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	connectDB := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))
	db, err := gorm.Open(mysql.Open(connectDB), &gorm.Config{})
	if err != nil {
		panic("Not Connect")
	} else {
		fmt.Println("Connected")
	}

	DB = db

	// Migrate model to DB

	// db.AutoMigrate(models.User{})
	// db.AutoMigrate(models.Konsul{})
	// db.AutoMigrate(models.Coupon{})
	// db.AutoMigrate(models.Book{})
	// db.AutoMigrate(models.InfoBlog{})
	// db.AutoMigrate(models.Ads{})
	// db.AutoMigrate(models.Course{})
	// db.AutoMigrate(models.TransactionBook{})
}
