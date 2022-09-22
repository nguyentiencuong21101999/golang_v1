package service

import (
	"main/src/database"
	"main/src/helpers/errors"
	UserDTO "main/src/modules/user/dtos"
	"main/src/modules/user/models"

	"github.com/mitchellh/mapstructure"
)

func SignIn(params *UserDTO.SignInDTO) (*UserDTO.UserDTO, *errors.ErrorsResp) {

	user := new(UserDTO.UserDTO)

	query := models.User{
		UserName: params.Username,
	}

	res := database.DB.First(&models.User{}, query)

	mapstructure.Decode(res.Value, &user)
	if res.RowsAffected == 0 && user.UserId == 0 {
		return nil, errors.UserNotFound
	} else {

		return user, nil
	}
}
