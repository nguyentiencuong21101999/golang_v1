package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"main/src/modules/user/models"
)

var DB *gorm.DB

func ConnectDatabase() {
	db, err := gorm.Open("mysql", "root:tiencuong@123@tcp(localhost:3306)/test")

	if err != nil {
		panic("Failed to connect to database!")
	}

	db.AutoMigrate(models.User{})

	DB = db
}
