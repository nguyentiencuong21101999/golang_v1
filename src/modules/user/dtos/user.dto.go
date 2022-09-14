package UserDTO

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type UserDTO struct {
	UserId   int32  `json:"userId"`
	UserName string `json:"userName"`
	Password string `json:"password"`
}

type SignInDTO struct {
	Username string `json:"userName" binding:"required,min=2,max=32"`
	Password string `json:"password" binding:"required"`
}

func FromReq(c *gin.Context) (*SignInDTO, error) {
	var params *SignInDTO
	err := c.ShouldBindBodyWith(&params, binding.JSON)
	return params, err
}
