package stringx

import "strings"

// ToTitle returns a copy of the string s with all Unicode letters that begin
func ToTitle(s string) string {
	return strings.ToTitle(s)
}
