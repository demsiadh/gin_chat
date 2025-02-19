package common

const (
	SUCCESS = 0
	FAIL    = -1
)

// Response 响应对象
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"msg"`
	Data    interface{} `json:"data,omitempty"`
}

// NewResponse 创建一个响应对象
func NewResponse(code int, message string, data interface{}) *Response {
	return &Response{
		Code:    code,
		Message: message,
		Data:    data,
	}
}

// NewSuccessResponseWithData 创建一个成功的响应对象
func NewSuccessResponseWithData(data interface{}) *Response {
	return NewResponse(SUCCESS, "success", data)
}

// NewErrorResponse 创建一个失败的响应对象
func NewErrorResponse(message string) *Response {
	return NewResponse(FAIL, message, nil)
}

// NewSuccessResponseWithMsg 创建一个成功的响应对象
func NewSuccessResponseWithMsg(msg string) *Response {
	return NewResponse(SUCCESS, msg, nil)
}

func NewSuccessResponse() *Response {
	return NewResponse(SUCCESS, "success", nil)
}
