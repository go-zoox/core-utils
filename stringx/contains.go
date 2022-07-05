package stringx

import "strings"

// Contains returns true if s contains substr.
func Contains(s, substr string) bool {
	return strings.Contains(s, substr)
}
