package main

import (
	"main/src/config"
	"main/src/database"
	console "main/src/helpers/consoles"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"

	UserController "main/src/modules/user/controller"
	UserMiddleware "main/src/modules/user/middleware"
)

type A struct {
	User string `json:"user"`
	Pas  string `json:"pas"`
}

func Test[T any](a []T) {
	console.Log(a)

}

func test(params A) {
	params.User = "a"

	console.Log(params)
}
func main() {

	router := gin.Default()

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		config.LoadConf()
		database.ConnectDatabase()

		router.GET("/heathcheck", func(c *gin.Context) {
			c.JSON(http.StatusOK, "heathcheck")
		})
		router.POST("/users/sign-in", UserMiddleware.TransformAndValidateSignInReq, UserController.SignIn)
		wg.Done()
	}()
	wg.Wait()
	router.Run(":4000")

}
