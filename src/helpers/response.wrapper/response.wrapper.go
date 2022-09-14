package responseWrapper

type ResponseWrapper struct {
	Data  interface{} `json:"data" `
	Error interface{} `json:"error" default:"nil"`
}

func New(data interface{}, err interface{}) *ResponseWrapper {
	response := &ResponseWrapper{
		Data:  data,
		Error: err,
	}
	return response
}
