package main

import (
	"main/src/database"
	console "main/src/helpers/consoles"
	"net/http"

	"github.com/gin-gonic/gin"

	UserController "main/src/modules/user/controllers"
)

type A struct {
	user string
	pas  []string
}

func main() {
	router := gin.Default()

	router.POST("/users/sign-in", UserController.SignIn)

	router.GET("/test", func(c *gin.Context) {
		a := A{
			user: "abc",
			pas: []string{
				"a",
			},
		}
		console.Log(a)
		c.JSON(http.StatusOK, gin.H{
			"data": a,
		})
	})
	database.ConnectDatabase()
	router.Run(":4000")
}
