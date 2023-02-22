package strings

import (
	"testing"

	"github.com/go-zoox/testify"
)

func TestTemplate(t *testing.T) {
	s, err := Template("hi, {{.name}}", map[string]any{
		"name": "Zero",
	})
	testify.Assert(t, err == nil)
	testify.Equal(t, s, "hi, Zero")
}
