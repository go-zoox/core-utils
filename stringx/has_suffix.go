package strings

import gostrings "strings"

// HasSuffix returns true if s ends with suffix.
func HasSuffix(s, suffix string) bool {
	return gostrings.HasSuffix(s, suffix)
}
