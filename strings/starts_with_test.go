package strings

import (
	"testing"

	"github.com/go-zoox/testify"
)

func TestStartsWith(t *testing.T) {
	msg := "@Zero hhh"
	// prefix := "@Zero "
	prefix := "@Zero\u2005"
	// fmt.Println(len(prefix), msg[0:len(prefix)], prefix == msg[0:len(prefix)])
	testify.Assert(t, StartsWith(msg, prefix))
}
