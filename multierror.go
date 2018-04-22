package errors

import (
	"fmt"
)

// Error 是一个可以跟踪多个错误的error 类型
type Error struct {
	Errors      []error
	ErrorFormat ErrorFormatFunc
}

func (e *Error) Error() string {
	fn := e.ErrorFormat
	if fn == nil {
		fn = ListFormatFunc
	}

	return fn(e.Errors)
}

func (e *Error) ErrorOrNil() error {
	if e == nil {
		return nil
	}

	if len(e.Errors) == 0 {
		return nil
	}

	return e
}

func (e *Error) String() string {
	return fmt.Sprintf("*%#v", *e)
}

func (e *Error) WrappedErrors() []error {
	return e.Errors
}
