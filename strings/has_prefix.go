package strings

import gostrings "strings"

// HasPrefix returns true if s starts with prefix.
func HasPrefix(s, prefix string) bool {
	return gostrings.HasPrefix(s, prefix)
}
