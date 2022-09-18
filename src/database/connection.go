package database

import (
	console "main/src/helpers/consoles"
	"main/src/modules/user/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func ConnectDatabase() {
	db, err := gorm.Open("mysql", "root:AbCd@1234@tcp(localhost:3306)/test")
	db.LogMode(true)
	if err != nil {
		panic("Failed to connect to database!")
	}

	db.AutoMigrate(models.User{})

	DB = db
	console.Log("Database connect successfully")
}
