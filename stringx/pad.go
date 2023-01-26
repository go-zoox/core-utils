package strings

import gostrings "strings"

// PadLeft ...
func PadLeft(s string, n int, padding string) string {
	return gostrings.Repeat(padding, n-len(s)) + s
}

// PadRight ...
func PadRight(s string, n int, padding string) string {
	return s + gostrings.Repeat(padding, n-len(s))
}

// PadZero ...
func PadZero(s string, n int) string {
	return PadRight(s, n, "0")
}
