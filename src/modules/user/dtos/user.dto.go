package dto

type UserDTO struct {
	UserId   int32  `json:"userId"`
	UserName string `json:"userName"`
	Password string `json:"password"`
}

type SignInDTO struct {
	Username string `json:"userName" binding:"required,min=1"`
	Password string `json:"password"`
}
