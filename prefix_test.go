package errors

import (
	"testing"

	"github.com/gotoxu/assert"
)

func TestPrefixError(t *testing.T) {
	original := &Error{
		Errors: []error{New("foo")},
	}

	result := Prefix(original, "bar")
	assert.DeepEqual(t, result.(*Error).Errors[0].Error(), "bar foo")
}

func TestPrefixNilError(t *testing.T) {
	var err error
	result := Prefix(err, "bar")
	assert.Nil(t, result)
}

func TestPrefixNonError(t *testing.T) {
	original := New("foo")
	result := Prefix(original, "bar")
	assert.DeepEqual(t, result.Error(), "bar foo")
}
