package strings

import "strconv"

// ToInt converts a string to an integer.
func ToInt(s string) (int, error) {
	return strconv.Atoi(s)
}

// MustToInt converts a string to an integer.
func MustToInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

// ToInt64 converts a string to an integer.
func ToInt64(s string) (int64, error) {
	return strconv.ParseInt(s, 10, 64)
}

// MustToInt64 converts a string to an integer.
func MustToInt64(s string) int64 {
	i, _ := strconv.ParseInt(s, 10, 64)
	return i
}

// ToBool converts a string to a boolean.
// It accepts 1, t, T, TRUE, true, True, 0, f, F, FALSE, false, False. Any other value returns an error.
func ToBool(s string) (bool, error) {
	return strconv.ParseBool(s)
}

// MustToBool converts a string to a boolean.
// It accepts 1, t, T, TRUE, true, True, 0, f, F, FALSE, false, False.
func MustToBool(s string) bool {
	b, _ := strconv.ParseBool(s)
	return b
}

// ToFloat converts a string to a float64.
func ToFloat(s string) (float64, error) {
	return strconv.ParseFloat(s, 64)
}

// MustToFloat converts a string to a float64.
func MustToFloat(s string) float64 {
	f, _ := strconv.ParseFloat(s, 64)
	return f
}

// ToFloat32 converts a string to a float32.
func ToFloat32(s string) (float32, error) {
	f, err := strconv.ParseFloat(s, 32)
	return float32(f), err
}

// MustToFloat32 converts a string to a float32.
func MustToFloat32(s string) float32 {
	f, _ := strconv.ParseFloat(s, 32)
	return float32(f)
}

// ToFloat64 converts a string to a float64.
func ToFloat64(s string) (float64, error) {
	return strconv.ParseFloat(s, 64)
}

// MustToFloat64 converts a string to a float64.
func MustToFloat64(s string) float64 {
	f, _ := strconv.ParseFloat(s, 64)
	return f
}
