package strings

import (
	"regexp"
	gostrings "strings"
)

// Replace returns a copy of s with the first n non-overlapping instances of old replaced by new.
// If old is empty, it matches at the beginning of the string and after each UTF-8 sequence, yielding up to k+1 replacements for a k-rune string.
// If n < 0, there is no limit on the number of replacements.
func Replace(s, old, new string, n int) string {
	return gostrings.Replace(s, old, new, n)
}

// ReplaceAll returns a copy of s with all non-overlapping instances of old replaced by new.
// If old is empty, it matches at the beginning of the string and after each UTF-8 sequence, yielding up to k+1 replacements for a k-rune string.
func ReplaceAll(s, old, new string) string {
	return gostrings.ReplaceAll(s, old, new)
}

// ReplaceAllFunc returns a copy of s with all non-overlapping instances of the function f applied to old replaced by new.
// If old is empty, it matches at the beginning of the string and after each UTF-8 sequence, yielding up to k+1 replacements for a k-rune string.
func ReplaceAllFunc(s string, re string, f func([]byte) []byte) string {
	return string(regexp.MustCompile(re).ReplaceAllFunc([]byte(s), f))
}
