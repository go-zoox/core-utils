package strings

import (
	"time"

	"github.com/go-zoox/core-utils/cast"
)

// Value is the value string to spfcast to new type with spf13/
type Value string

// Int ...
func (s Value) Int() int {
	return cast.ToInt(s.String())
}

// Int64 ...
func (s Value) Int64() int64 {
	return cast.ToInt64(s.String())
}

// UInt ...
func (s Value) UInt() uint {
	return cast.ToUint(s.String())
}

// Uint64 ...
func (s Value) Uint64() uint64 {
	return cast.ToUint64(s.String())
}

// Bool ...
func (s Value) Bool() bool {
	return cast.ToBool(s.String())
}

// Float64 ...
func (s Value) Float64() float64 {
	return cast.ToFloat64(s.String())
}

// Time ...
func (s Value) Time() time.Time {
	return cast.ToTime(s.String())
}

// Duration ...
func (s Value) Duration() time.Duration {
	return cast.ToDuration(s.String())
}

// String ...
func (s Value) String() string {
	return string(s)
}
