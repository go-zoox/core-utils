package cast

import (
	"time"

	spfcast "github.com/spf13/cast"
)

// ToInt casts a interfacfe to int.
func ToInt(i interface{}) int {
	return spfcast.ToInt(i)
}

// ToInt64 casts a interfacfe to int64.
func ToInt64(i interface{}) int64 {
	return spfcast.ToInt64(i)
}

// ToUint casts a interfacfe to uint.
func ToUint(i interface{}) uint {
	return spfcast.ToUint(i)
}

// ToUint64 casts a interfacfe to uint64.
func ToUint64(i interface{}) uint64 {
	return spfcast.ToUint64(i)
}

// ToBool casts a interfacfe to bool.
// It accepts
//
//	true: 1, t, T, TRUE, true, True
//	false: 0, f, F, FALSE, false, False, any other value.
func ToBool(i interface{}) bool {
	return spfcast.ToBool(i)
}

// ToFloat64 casts a interfacfe to float64.
func ToFloat64(i interface{}) float64 {
	return spfcast.ToFloat64(i)
}

// ToTime casts a interfacfe to time.Time.
func ToTime(i interface{}) time.Time {
	return spfcast.ToTime(i)
}

// ToDuration casts a interfacfe to time.Duration.
func ToDuration(i interface{}) time.Duration {
	return spfcast.ToDuration(i)
}

// ToString casts a interface to string.
func ToString(i interface{}) string {
	return spfcast.ToString(i)
}
