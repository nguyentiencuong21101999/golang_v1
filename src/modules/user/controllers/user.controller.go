package controller

import (
	"main/src/database"
	console "main/src/helpers/consoles"
	"main/src/modules/user/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SignIn(c *gin.Context) {
	//params := c.MustGet("params")

	result := database.DB.First(&models.User{UserId: 1})

	console.Log(result.Value)

	c.JSON(http.StatusOK, gin.H{
		"data": result,
	})

}
