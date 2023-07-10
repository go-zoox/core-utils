package reflect

import (
	"testing"

	"github.com/go-zoox/testify"
)

func TestIsFunc(t *testing.T) {
	testify.AssertFalse(t, IsFunc(nil))

	fn := func() {}
	testify.Assert(t, IsFunc(fn))
}
