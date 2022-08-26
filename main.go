package main

import (
	"main/src/database"

	"github.com/gin-gonic/gin"

	controller "main/src/modules/user/controllers"
	middleware "main/src/modules/user/middlewares"
)

func main() {
	router := gin.Default()
	router.POST("/users/sign-in",middleware.TransformAndValidateSignInReq, controller.SignIn)

	database.ConnectDatabase()
	router.Run(":4000")
}
