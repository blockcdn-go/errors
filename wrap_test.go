package errors

import (
	"fmt"
	"testing"

	"github.com/gotoxu/assert"
)

func TestGetAll(t *testing.T) {
	cases := []struct {
		Err error
		Msg string
		Len int
	}{
		{},
		{
			fmt.Errorf("foo"),
			"foo",
			1,
		},
		{
			fmt.Errorf("bar"),
			"foo",
			0,
		},
		{
			Wrapf("bar", fmt.Errorf("foo")),
			"foo",
			1,
		},
		{
			Wrapf("{{err}}", fmt.Errorf("foo")),
			"foo",
			2,
		},
		{
			Wrapf("bar", Wrapf("baz", fmt.Errorf("foo"))),
			"foo",
			1,
		},
	}

	for _, tc := range cases {
		actual := GetAll(tc.Err, tc.Msg)
		assert.Len(t, actual, tc.Len)

		for _, v := range actual {
			assert.DeepEqual(t, v.Error(), tc.Msg)
		}
	}
}

func TestGetAllType(t *testing.T) {
	cases := []struct {
		Err  error
		Type interface{}
		Len  int
	}{
		{},
		{
			fmt.Errorf("foo"),
			"foo",
			0,
		},
		{
			fmt.Errorf("bar"),
			fmt.Errorf("foo"),
			1,
		},
		{
			Wrapf("bar", fmt.Errorf("foo")),
			fmt.Errorf("baz"),
			2,
		},
		{
			Wrapf("bar", Wrapf("baz", fmt.Errorf("foo"))),
			Wrapf("", nil),
			0,
		},
	}

	for _, tc := range cases {
		actual := GetAllType(tc.Err, tc.Type)
		assert.Len(t, actual, tc.Len)
	}
}
