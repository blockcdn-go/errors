package errors

import (
	"testing"

	"github.com/gotoxu/assert"
)

func TestListFormatFuncSingle(t *testing.T) {
	expected := `1 error occurred: 

* foo`

	errors := []error{
		New("foo"),
	}

	actual := ListFormatFunc(errors)
	assert.DeepEqual(t, actual, expected)
}

func TestListFormatFuncMultiple(t *testing.T) {
	expected := `2 errors occurred: 

* foo
* bar`

	errors := []error{
		New("foo"),
		New("bar"),
	}

	actual := ListFormatFunc(errors)
	assert.DeepEqual(t, actual, expected)
}
