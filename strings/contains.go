package strings

import gostrings "strings"

// Contains returns true if s contains substr.
func Contains(s, substr string) bool {
	return gostrings.Contains(s, substr)
}
