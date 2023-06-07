package typ

import (
	"testing"

	"github.com/go-zoox/testify"
)

func TestValue(t *testing.T) {
	value := map[string]interface{}{
		"foo":  "bar",
		"foo5": 1,
		"foo6": 1.1,
		"foo7": true,
		"foo8": nil,
		"foo2": map[string]interface{}{
			"foo3": "bar3",
		},
		"foo4": []string{
			"bar4",
			"bar5",
		},
	}

	testify.Equal(t, "bar", NewValue(value).Get("foo").String())
	testify.Equal(t, 1, NewValue(value).Get("foo5").Int())
	testify.Equal(t, 1.1, NewValue(value).Get("foo6").Float64())
	testify.Equal(t, true, NewValue(value).Get("foo7").Bool())
	testify.Equal(t, false, NewValue(value).Get("foo8").Bool())
	testify.Equal(t, false, NewValue(value).Get("foo9").Bool())
	testify.Equal(t, false, NewValue(value).Get("foo6").Bool())

	// key path
	testify.Equal(t, "bar3", NewValue(value).Get("foo2.foo3").String())
	testify.Equal(t, "bar4", NewValue(value).Get("foo4.0").String())
	testify.Equal(t, "bar5", NewValue(value).Get("foo4.1").String())
}
