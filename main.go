package main

import (
	"main/src/database"

	"github.com/gin-gonic/gin"

	controllers "main/src/modules/user"
)

func main() {
	router := gin.Default()
	router.POST("/users/sign-in", controllers.SignIn)

	database.ConnectDatabase()
	router.Run(":4000")
}
