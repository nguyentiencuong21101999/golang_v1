package middleware

import (
	UserDTO "main/src/modules/user/dtos"
	"net/http"

	"github.com/gin-gonic/gin"
)

func TransformAndValidateSignInReq(c *gin.Context) {
	var params *UserDTO.SignInDTO
	err := c.ShouldBindJSON(&params)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
		c.Abort()
	}
	c.Set("params", params)
}
