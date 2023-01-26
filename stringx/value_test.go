package stringx

import (
	"testing"
	"time"

	"github.com/go-zoox/testify"
)

func TestValue(t *testing.T) {
	v := Value("1000")

	testify.Equal(t, int(1000), v.ToInt())
	testify.Equal(t, 1000, v.ToInt64())
	testify.Equal(t, 1000, v.ToUInt())
	testify.Equal(t, 1000, v.ToUint64())
	testify.Equal(t, 1000, v.ToFloat64())
	testify.Equal(t, false, v.ToBool())
	testify.Equal(t, time.Duration(1000), v.ToDuration())
}
