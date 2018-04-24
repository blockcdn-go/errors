package errors

import (
	"fmt"
	"net/http"
)

// HTTPError 用于封装HTTP请求的相关错误
type HTTPError struct {
	status int
	text   string
	err    error
}

// Status 返回HTTPError中包含的状态码
func (e *HTTPError) Status() int {
	return e.status
}

func (e *HTTPError) Error() string {
	if e.err == nil {
		return e.text
	}

	return e.err.Error()
}

func (e *HTTPError) String() string {
	return fmt.Sprintf("*%#v", *e)
}

// NewHTTP 创建一个新的HTTPError
func NewHTTP(code int, err error) *HTTPError {
	return &HTTPError{
		status: code,
		text:   http.StatusText(code),
		err:    err,
	}
}
