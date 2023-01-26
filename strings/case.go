package strings

import (
	gostrings "strings"
)

// UpperCase returns a copy of s with all lowercase letters converted to
// uppercase.
func UpperCase(s string) string {
	return gostrings.ToUpper(s)
}

// LowerCase returns a copy of s with all uppercase letters converted to
// lowercase.
func LowerCase(s string) string {
	return gostrings.ToLower(s)
}

// ToUpper returns a copy of s with all lowercase letters converted to
// uppercase.
func ToUpper(s string) string {
	return gostrings.ToUpper(s)
}

// ToLower returns a copy of s with all uppercase letters converted to
// lowercase.
func ToLower(s string) string {
	return gostrings.ToLower(s)
}
