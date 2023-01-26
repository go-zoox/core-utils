package strings

import gostrings "strings"

// Trim returns a substring of s without leading and trailing cutsets.
func Trim(s, cutset string) string {
	return gostrings.Trim(s, cutset)
}

// TrimSpace returns a substring of s without leading and trailing spaces.
func TrimSpace(s string) string {
	return gostrings.TrimSpace(s)
}

// TrimLeft returns a substring of s without leading cutsets.
func TrimLeft(s, cutset string) string {
	return gostrings.TrimLeft(s, cutset)
}

// TrimRight returns a substring of s without trailing cutsets.
func TrimRight(s, cutset string) string {
	return gostrings.TrimRight(s, cutset)
}
