package controller

import (
	"main/src/helpers/errors"
	responseWrapper "main/src/helpers/response.wrapper"
	UserDTO "main/src/modules/user/dtos"
	"main/src/modules/user/service"

	"github.com/gin-gonic/gin"
)

func SignIn(c *gin.Context) {
	params, _ := UserDTO.FromReq(c)
	data, err := service.SignIn(params)
	if err != nil {
		errors.HandleError(err, c)
		return
	}
	responseWrapper.HandleResp(responseWrapper.New(data, nil), c)
}
