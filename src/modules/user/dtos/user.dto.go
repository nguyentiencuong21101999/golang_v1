package UserDTO

import (
	"net/http"

	"github.com/gin-gonic/gin"
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
	err := c.ShouldBindJSON(&params)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
		c.Abort()
	}
	return params, nil
}
