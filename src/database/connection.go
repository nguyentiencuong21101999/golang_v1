package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open("mysql", "root:tiencuong@123@tcp(localhost:3306)/wolfden")

	if err != nil {
		panic("Failed to connect to database!")
	}

	//   database.AutoMigrate(&Book{})

	DB = database
}
