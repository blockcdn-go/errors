package errors

import (
	"testing"

	"github.com/gotoxu/assert"
)

func TestAppendError(t *testing.T) {
	original := &Error{
		Errors: []error{New("foo")},
	}

	result := Append(original, New("bar"))
	assert.Len(t, result.Errors, 2)

	original = &Error{}
	result = Append(original, New("bar"))
	assert.Len(t, result.Errors, 1)

	var e *Error
	result = Append(e, New("bar"))
	assert.Len(t, result.Errors, 1)

	original = &Error{
		Errors: []error{New("foo")},
	}

	result = Append(original, Append(nil, New("foo")), New("bar"))
	assert.Len(t, result.Errors, 3)
}

func TestAppendNilError(t *testing.T) {
	var err error
	result := Append(err, New("bar"))
	assert.Len(t, result.Errors, 1)
}

func TestAppendNilErrorArgs(t *testing.T) {
	var err error
	var nilErr *Error
	result := Append(err, nilErr)
	assert.Len(t, result.Errors, 0)
}

func TestAppendNilErrorIfaceArg(t *testing.T) {
	var err error
	var nilErr error
	result := Append(err, nilErr)
	assert.Len(t, result.Errors, 0)
}

func TestAppendNonError(t *testing.T) {
	original := New("foo")
	result := Append(original, New("bar"))
	assert.Len(t, result.Errors, 2)
}

func TestAppendNonErrorError(t *testing.T) {
	original := New("foo")
	result := Append(original, Append(nil, New("bar")))
	assert.Len(t, result.Errors, 2)
}
