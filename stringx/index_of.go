package stringx

import "strings"

// IndexOf returns the index of the first occurrence of substr in s, or -1 if substr is not found.
func IndexOf(s, substr string) int {
	return strings.Index(s, substr)
}
