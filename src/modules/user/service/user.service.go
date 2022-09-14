package service

import (
	"main/src/database"
	console "main/src/helpers/consoles"
	UserDTO "main/src/modules/user/dtos"
	"main/src/modules/user/models"
)

func SignIn(params *UserDTO.SignInDTO) interface{} {
	res := database.DB.First(&models.User{})
	console.Log(res)
	return res
}
