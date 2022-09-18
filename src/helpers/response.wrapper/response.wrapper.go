package responseWrapper

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseWrapper struct {
	Data  interface{} `json:"data" `
	Error interface{} `json:"error" `
}

func New(data interface{}, err interface{}) *ResponseWrapper {
	response := &ResponseWrapper{
		Data:  data,
		Error: err,
	}
	return response
}

func HandleResp(data *ResponseWrapper, c *gin.Context) {
	c.JSON(http.StatusOK, data)
}
