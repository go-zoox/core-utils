package stringx

import "strings"

// Equals returns true if s1 and s2 are equal.
func Equals(s, t string) bool {
	return strings.EqualFold(s, t)
}
