package stringx

import "strings"

// HasPrefix returns true if s starts with prefix.
func HasPrefix(s, prefix string) bool {
	return strings.HasPrefix(s, prefix)
}
