package responseWrapper

type ResponseWrapper struct {
	data  interface{}
	error interface{}
}

func New(data interface{}) *ResponseWrapper {
	var res *ResponseWrapper
	res.data = data
	res.error = nil
	return res
}
