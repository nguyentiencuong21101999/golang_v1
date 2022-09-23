package middleware

import (
	"main/src/helpers/errors"
	UserDTO "main/src/modules/user/dtos"

	"github.com/gin-gonic/gin"
)

func TransformAndValidateSignInReq(c *gin.Context) {
	_, err := UserDTO.FromReq(c)
	if err != nil {
		errors.HandleError(errors.ParseError(err), c)
		c.Abort()
	}
	c.Next()
}
