package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"main/src/helpers/errors"
	userDTO "main/src/modules/user/dtos"
)

func TransformAndValidateSignInReq(c *gin.Context) {
	var params userDTO.SignInDTO
	if err := c.ShouldBindJSON(&params); err != nil {
		errs := new(errors.ErrorsResp)

		er := errs.Init("err.abc", "abc", 400)
		fmt.Println("abc", er)
		c.JSON(http.StatusOK, gin.H{
			"datas": c.ShouldBindJSON(&e),
		})

		// c.AbortWithError(err.Ini)
		return
	}
	return
	// c.Next()
}
