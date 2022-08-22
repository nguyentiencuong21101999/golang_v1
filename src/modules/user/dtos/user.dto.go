package UserDto

type UserDTO struct {
	UserId   int32  `json:"userId"`
	UserName string `json:"userName"`
	Password string `json:"password"`
}

type SignInDTO struct {
	Username string `json:"userName"`
	Password string `json:"password"`
}
