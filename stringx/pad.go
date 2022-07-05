package stringx

import "strings"

// PadLeft ...
func PadLeft(s string, n int, padding string) string {
	return strings.Repeat(padding, n) + s
}

// PadRight ...
func PadRight(s string, n int, padding string) string {
	return s + strings.Repeat(padding, n)
}

// PadZero ...
func PadZero(s string, n int) string {
	return PadRight(s, n, "0")
}
