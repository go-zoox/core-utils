package strings

import gostrings "strings"

// Split splits a string into a slice of strings.
func Split(s, sep string) []string {
	return gostrings.Split(s, sep)
}

// SplitN splits a string into a slice of strings.
func SplitN(s, sep string, n int) []string {
	return gostrings.SplitN(s, sep, n)
}
