package stringx

import "strings"

// HasSuffix returns true if s ends with suffix.
func HasSuffix(s, suffix string) bool {
	return strings.HasSuffix(s, suffix)
}
