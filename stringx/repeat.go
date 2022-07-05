package stringx

import "strings"

// Repeat returns a new string consisting of n copies of the string s.
func Repeat(s string, n int) string {
	return strings.Repeat(s, n)
}
