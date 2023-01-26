package cast

import (
	"testing"
	"time"

	"github.com/go-zoox/testify"
)

func TestCast(t *testing.T) {
	// int
	testify.Equal(t, 100, ToInt("100"))
	testify.Equal(t, int64(100), ToInt64("100"))
	testify.Equal(t, uint(100), ToUint("100"))
	testify.Equal(t, uint64(100), ToUint64("100"))

	// float
	testify.Equal(t, float64(100.00), ToFloat64("100.00"))

	// bool
	testify.Equal(t, true, ToBool("true"))
	testify.Equal(t, true, ToBool("True"))
	testify.Equal(t, true, ToBool("T"))
	testify.Equal(t, true, ToBool(1))
	testify.Equal(t, true, ToBool(-1))
	testify.Equal(t, false, ToBool("false"))
	testify.Equal(t, false, ToBool("False"))
	testify.Equal(t, false, ToBool("F"))
	testify.Equal(t, false, ToBool(0))
	// any other value return false
	testify.Equal(t, false, ToBool("Yes"))
	testify.Equal(t, false, ToBool("No"))
	//
	testify.Equal(t, false, ToBool(nil))

	// time
	now := time.Now()
	testify.Equal(t, now.Format("2006-01-02 15:04:05"), ToTime(now.Format("2006-01-02 15:04:05")).Format("2006-01-02 15:04:05"))
	testify.Equal(t, time.Duration(1), ToDuration(1))
}
