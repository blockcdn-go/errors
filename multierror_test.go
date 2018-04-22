package errors

import (
	"testing"

	"github.com/gotoxu/assert"
)

func TestErrorErrorCustom(t *testing.T) {
	errors := []error{
		New("foo"),
		New("bar"),
	}

	fn := func(es []error) string {
		return "foo"
	}

	multi := &Error{Errors: errors, ErrorFormat: fn}
	assert.DeepEqual(t, multi.Error(), "foo")
}

func TestErrorErrorDefault(t *testing.T) {
	expected := `2 errors occurred: 

* foo
* bar`

	errors := []error{
		New("foo"),
		New("bar"),
	}

	multi := &Error{Errors: errors}
	assert.DeepEqual(t, multi.Error(), expected)
}

func TestErrorErrorOrNil(t *testing.T) {
	err := new(Error)
	assert.Nil(t, err.ErrorOrNil())

	err.Errors = []error{New("foo")}
	v := err.ErrorOrNil()
	assert.NotNil(t, v)
	assert.DeepEqual(t, v, err)
}

func TestErrorWrappedErrors(t *testing.T) {
	errors := []error{
		New("foo"),
		New("bar"),
	}

	multi := &Error{Errors: errors}
	assert.DeepEqual(t, multi.Errors, multi.WrappedErrors())
}
