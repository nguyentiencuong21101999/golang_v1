package main

import (
	"main/src/config"
	"main/src/database"
	"net/http"

	"github.com/gin-gonic/gin"

	UserController "main/src/modules/user/controllers"
	UserMiddleware "main/src/modules/user/middlewares"
)

func main() {
	router := gin.Default()
	config.LoadConf()

	router.POST("/users/sign-in", UserMiddleware.TransformAndValidateSignInReq, UserController.SignIn)
	router.GET("/heathCheck", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"data": "success",
		})
	})

	config.GetConfig("dbUri")
	//console.Log(config.GetConfig("port"))
	database.ConnectDatabase()
	router.Run(":4000")
}
