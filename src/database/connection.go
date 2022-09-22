package database

import (
	"main/src/config"
	c "main/src/helpers/consoles"
	"main/src/modules/user/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func ConnectDatabase() {
	c.Log(config.GetConfig().DbUri)
	db, err := gorm.Open("mysql", config.GetConfig().DbUri)
	db.LogMode(true)
	if err != nil {
		panic("Failed to connect to database!")
	}

	db.AutoMigrate(models.User{})

	DB = db
	c.Log("Database connect successfully")
}
