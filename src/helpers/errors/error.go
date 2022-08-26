package errors

type ErrorsResp struct {
	status  int32
	code    string
	message string
}

type Errors struct {
	
}

func (err *ErrorsResp) Init(code string, message string, status int32) *ErrorsResp {
	err.code = code
	err.message = message
	err.status = status
	return err
}

// const errros = []struct{
// 	a string
// }{
// 	"a"
// 		//  new(ErrorsResp).Init("error.badRequest", "Bad request", 400)
	
// }
