package strings

import (
	"time"

	"github.com/go-zoox/core-utils/cast"
)

// Value is the value string to spfcast to new type with spf13/
type Value string

// ToInt ...
func (s Value) ToInt() int {
	return cast.ToInt(s.ToString())
}

// ToInt64 ...
func (s Value) ToInt64() int64 {
	return cast.ToInt64(s.ToString())
}

// ToUInt ...
func (s Value) ToUInt() uint {
	return cast.ToUint(s.ToString())
}

// ToUint64 ...
func (s Value) ToUint64() uint64 {
	return cast.ToUint64(s.ToString())
}

// ToBool ...
func (s Value) ToBool() bool {
	return cast.ToBool(s.ToString())
}

// ToFloat64 ...
func (s Value) ToFloat64() float64 {
	return cast.ToFloat64(s.ToString())
}

// ToTime ...
func (s Value) ToTime() time.Time {
	return cast.ToTime(s.ToString())
}

// ToDuration ...
func (s Value) ToDuration() time.Duration {
	return cast.ToDuration(s.ToString())
}

// ToString ...
func (s Value) ToString() string {
	return string(s)
}
