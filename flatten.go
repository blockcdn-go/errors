package errors

// Flatten 讲任意个*Error合并到一个单独的*Error中
func Flatten(err error) error {
	if _, ok := err.(*Error); !ok {
		return err
	}

	flatErr := new(Error)
	flatten(err, flatErr)
	return flatErr
}

func flatten(err error, flatErr *Error) {
	switch err := err.(type) {
	case *Error:
		for _, e := range err.Errors {
			flatten(e, flatErr)
		}
	default:
		flatErr.Errors = append(flatErr.Errors, err)
	}
}
