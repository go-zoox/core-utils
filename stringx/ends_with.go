package stringx

import "strings"

// EndsWith returns true if s ends with suffix.
func EndsWith(s, suffix string) bool {
	return strings.HasSuffix(s, suffix)
}
