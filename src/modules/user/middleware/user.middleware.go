package middleware

import (
	"main/src/helpers/errors"
	UserDTO "main/src/modules/user/dtos"

	"github.com/gin-gonic/gin"
)

func TransformAndValidateSignInReq(c *gin.Context) {
	_, err := UserDTO.FromReq(c)
	if err != nil {
		errors.HandleError(errors.BadRequest, c)
		// c.JSON(http.StatusOK, gin.H{
		// 	"error": err.Error(),
		// })
		c.Abort()
	}
	c.Next()
}
