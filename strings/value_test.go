package strings

import (
	"testing"
	"time"

	"github.com/go-zoox/testify"
)

func TestValue(t *testing.T) {
	v := Value("1000")

	testify.Equal(t, int(1000), v.Int())
	testify.Equal(t, 1000, v.Int64())
	testify.Equal(t, 1000, v.UInt())
	testify.Equal(t, 1000, v.Uint64())
	testify.Equal(t, 1000, v.Float64())
	testify.Equal(t, false, v.Bool())
	testify.Equal(t, time.Duration(1000), v.Duration())
}
