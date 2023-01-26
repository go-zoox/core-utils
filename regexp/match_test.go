package regexp

import (
	"testing"

	"github.com/go-zoox/testify"
)

func TestMatch(t *testing.T) {
	testify.Equal(t, false, Match("^\\s+", "123"))
	testify.Equal(t, true, Match("^\\d+", "123"))
	testify.Equal(t, true, Match("^\\s+", " 123"))
}
