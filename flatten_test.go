package errors

import (
	"fmt"
	"strings"
	"testing"

	"github.com/gotoxu/assert"
)

func TestFlatten(t *testing.T) {
	original := &Error{
		Errors: []error{
			New("one"),
			&Error{
				Errors: []error{
					New("two"),
					&Error{
						Errors: []error{
							New("three"),
						},
					},
				},
			},
		},
	}

	expected := strings.TrimSpace(`
3 errors occurred: 

* one
* two
* three
		`)

	actual := fmt.Sprintf("%s", Flatten(original))
	assert.DeepEqual(t, actual, expected)
}

func TestFlattenNonError(t *testing.T) {
	err := New("foo")
	actual := Flatten(err)
	assert.DeepEqual(t, actual, err)
}
