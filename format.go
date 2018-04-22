package errors

import (
	"fmt"
	"strings"
)

// ErrorFormatFunc 是一个回调函数，可以讲一个error的数组转换为一个字符串
type ErrorFormatFunc func([]error) string

// ListFormatFunc 是一个基本的错误数组格式化函数
func ListFormatFunc(es []error) string {
	if len(es) == 1 {
		return fmt.Sprintf("1 error occurred: \n\n* %s", es[0])
	}

	points := make([]string, len(es))
	for i, err := range es {
		points[i] = fmt.Sprintf("* %s", err)
	}

	return fmt.Sprintf("%d errors occurred: \n\n%s", len(es), strings.Join(points, "\n"))
}
