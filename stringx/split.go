package stringx

import "strings"

// Split splits a string into a slice of strings.
func Split(s, sep string) []string {
	return strings.Split(s, sep)
}

// SplitN splits a string into a slice of strings.
func SplitN(s, sep string, n int) []string {
	return strings.SplitN(s, sep, n)
}
