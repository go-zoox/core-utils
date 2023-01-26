package strings

import gostrings "strings"

// StartsWith returns true if s starts with prefix.
func StartsWith(s, prefix string) bool {
	return gostrings.HasPrefix(s, prefix)
}
