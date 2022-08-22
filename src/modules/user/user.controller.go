package controllers

import (
	"log"
	"net/http"

	UserDto "main/src/modules/user/dtos"

	"github.com/gin-gonic/gin"
)

func SignIn(c *gin.Context) {
	var params UserDto.SignInDTO
	// if err := c.ShouldBindBodyWith(&params, binding.JSON); err != nil {
	// 	log.Printf("%+v", err)
	// }
	if err := c.ShouldBindJSON(&params); err != nil {
		log.Printf("%+v", err)
	}
	c.JSON(http.StatusOK, gin.H{
		"data": params,
	})

}
