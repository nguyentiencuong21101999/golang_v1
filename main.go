package main

import (
	"net/http"

	"main/src/database"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello world"})
	})

	database.ConnectDatabase()
	r.Run()
}
