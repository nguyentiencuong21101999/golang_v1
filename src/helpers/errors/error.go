package errors

import (
	responseWrapper "main/src/helpers/response.wrapper"

	"github.com/gin-gonic/gin"
)

type ErrorsResp struct {
	Status  int    `json:"status,omitempty" `
	Code    string `json:"code"`
	Message string `json:"message"`
}

func NewError(code string, message string, status int) *ErrorsResp {
	err := &ErrorsResp{
		Code:    code,
		Message: message,
		Status:  status,
	}
	return err
}

var (
	BadRequest          *ErrorsResp = NewError("error.badRequest", "Bad request", 400)
	Unauthorized        *ErrorsResp = NewError("error.unauthorized", "Unauthorized", 401)
	Forbidden           *ErrorsResp = NewError("error.forbiden", "Forbidden", 403)
	Sensitive           *ErrorsResp = NewError("error.sensitive", "An error occurred, please try again later.", 400)
	InternalServerError *ErrorsResp = NewError("error.internalServerError", "Internal server error.", 500)
	AccountBlocked      *ErrorsResp = NewError("error.accountBlocked", "Account is blocked.", 400)
	UserNotFound        *ErrorsResp = NewError("error.userNotFound", "User not found.", 400)
)

func ParseError(err error) *ErrorsResp {
	parseError := &ErrorsResp{
		Code:    InternalServerError.Code,
		Message: err.Error(),
		Status:  InternalServerError.Status,
	}
	return parseError
}

func HandleError(err *ErrorsResp, c *gin.Context) {
	c.JSON(err.Status, responseWrapper.New(nil, &ErrorsResp{
		Code:    err.Code,
		Message: err.Message,
	}))
	c.Abort()

}

// func HandleError(err *ErrorsResp, c *gin.Context) {
// 	if err.Status != http.StatusInternalServerError {
// 		c.JSON(err.Status, responseWrapper.New(nil, &Errors{
// 			Code:    err.Code,
// 			message: err.Message,
// 		}))
// 	} else {
// 		c.JSON(err.Status, responseWrapper.New(nil, &Errors{
// 			Code:    err.Code,
// 			message: err.Message,
// 		}))
// 	}

// }
