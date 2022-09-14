package main

import (
	"main/src/config"
	"main/src/database"
	console "main/src/helpers/consoles"
	responseWrapper "main/src/helpers/response.wrapper"
	"net/http"

	"github.com/gin-gonic/gin"

	UserController "main/src/modules/user/controller"
	UserMiddleware "main/src/modules/user/middleware"
)

type A struct {
	User string `json:"user"`
	Pas  int    `json:"pas"`
}

func Test[T any](a []T) {
	console.Log(a)

}
func main() {
	// Test([]A{
	// 	{User: "a", Pas: 123},
	// })
	router := gin.Default()
	config.LoadConf()

	router.POST("/users/sign-in", UserMiddleware.TransformAndValidateSignInReq, UserController.SignIn)
	router.GET("/heathCheck", func(c *gin.Context) {
		a := A{
			User: "abc",
			Pas:  1,
		}
		c.JSON(http.StatusOK, responseWrapper.New(a, nil))
	})

	// config.GetConfig("dbUri")
	//console.Log(config.GetConfig("port"))
	database.ConnectDatabase()
	router.Run(":4000")

}
