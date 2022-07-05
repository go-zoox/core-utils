package stringx

import "strings"

// Trim returns a substring of s without leading and trailing cutsets.
func Trim(s, cutset string) string {
	return strings.Trim(s, cutset)
}

// TrimSpace returns a substring of s without leading and trailing spaces.
func TrimSpace(s string) string {
	return strings.TrimSpace(s)
}

// TrimLeft returns a substring of s without leading cutsets.
func TrimLeft(s, cutset string) string {
	return strings.TrimLeft(s, cutset)
}

// TrimRight returns a substring of s without trailing cutsets.
func TrimRight(s, cutset string) string {
	return strings.TrimRight(s, cutset)
}
