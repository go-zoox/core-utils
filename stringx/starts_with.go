package stringx

import "strings"

// StartsWith returns true if s starts with prefix.
func StartsWith(s, prefix string) bool {
	return strings.HasPrefix(s, prefix)
}
