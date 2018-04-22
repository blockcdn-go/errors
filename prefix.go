package errors

import (
	"fmt"
)

// Prefix 为给定error添加一些前缀文字
func Prefix(err error, prefix string) error {
	if err == nil {
		return nil
	}

	format := fmt.Sprintf("%s {{err}}", prefix)
	switch err := err.(type) {
	case *Error:
		if err == nil {
			err = new(Error)
		}

		for i, e := range err.Errors {
			err.Errors[i] = Wrapf(format, e)
		}

		return err
	default:
		return Wrapf(format, err)
	}
}
