package controller

import (
	"net/http"

	responseWrapper "main/src/helpers/response.wrapper"
	UserDTO "main/src/modules/user/dtos"
	"main/src/modules/user/service"

	"github.com/gin-gonic/gin"
)

func SignIn(c *gin.Context) {
	params, _ := UserDTO.FromReq(c)
	res := service.SignIn(params)
	c.JSON(http.StatusOK, responseWrapper.New(&res, nil))

}
