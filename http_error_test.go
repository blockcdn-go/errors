package errors

import (
	"net/http"
	"testing"

	"github.com/gotoxu/assert"
)

func TestHTTPError(t *testing.T) {
	e := NewHTTP(400, New("foo"))
	assert.DeepEqual(t, e.Status(), http.StatusBadRequest)
}

func TestHTTPErrorError(t *testing.T) {
	e := NewHTTP(400, New("bar"))
	assert.DeepEqual(t, e.Error(), "bar")
}
