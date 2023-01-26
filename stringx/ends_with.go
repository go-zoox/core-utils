package strings

import gostrings "strings"

// EndsWith returns true if s ends with suffix.
func EndsWith(s, suffix string) bool {
	return gostrings.HasSuffix(s, suffix)
}
