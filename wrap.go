package errors

import (
	"reflect"
	"strings"
)

// WalkFunc 用作Walk方法的回调函数参数
type WalkFunc func(error)

// Wrapper 是一个error包装器接口
type Wrapper interface {
	WrappedErrors() []error
}

// Wrap 定义了一个错误包装器，并返回error
// 其中outer的message将作为返回的error的message
func Wrap(outer, inner error) error {
	return &wrappedError{
		Outer: outer,
		Inner: inner,
	}
}

// Wrapf 使用一个格式化的字符串包装一个error
func Wrapf(format string, err error) error {
	outerMsg := "<nil>"
	if err != nil {
		outerMsg = err.Error()
	}

	outer := New(strings.Replace(format, "{{err}}", outerMsg, -1))
	return Wrap(outer, err)
}

// Contains 判断给订单err中是否包含错误信息与msg相同的error
func Contains(err error, msg string) bool {
	return len(GetAll(err, msg)) > 0
}

// ContainsType 判断给定的err中是否包含与v相同类型的error
func ContainsType(err error, v interface{}) bool {
	return len(GetAllType(err, v)) > 0
}

// Get 与 GetAll 功能相同，只是返回最内层的error
func Get(err error, msg string) error {
	es := GetAll(err, msg)
	if len(es) > 0 {
		return es[len(es)-1]
	}
	return nil
}

// GetType 与 GetAllType的功能相同，只是返回最内层的error
func GetType(err error, v interface{}) error {
	es := GetAllType(err, v)
	if len(es) > 0 {
		return es[len(es)-1]
	}

	return nil
}

// GetAll 获取所有错误消息与给定msg相同，并被包装在err中的error
func GetAll(err error, msg string) []error {
	var result []error

	Walk(err, func(err error) {
		if err.Error() == msg {
			result = append(result, err)
		}
	})

	return result
}

// GetAllType 获取所有与给定v类型相同的error
func GetAllType(err error, v interface{}) []error {
	var result []error

	var search string
	if v != nil {
		search = reflect.TypeOf(v).String()
	}
	Walk(err, func(err error) {
		var needle string
		if err != nil {
			needle = reflect.TypeOf(err).String()
		}

		if needle == search {
			result = append(result, err)
		}
	})

	return result
}

// Walk 访问所有被包装的error并调用回调函数
func Walk(err error, cb WalkFunc) {
	if err == nil {
		return
	}

	switch e := err.(type) {
	case *wrappedError:
		cb(e.Outer)
		Walk(e.Inner, cb)
	case Wrapper:
		cb(err)
		for _, err := range e.WrappedErrors() {
			Walk(err, cb)
		}
	default:
		cb(err)
	}
}

type wrappedError struct {
	Outer error
	Inner error
}

func (w *wrappedError) Error() string {
	return w.Outer.Error()
}

func (w *wrappedError) WrappedErrors() []error {
	return []error{w.Outer, w.Inner}
}
