package strings

import gostrings "strings"

// ToTitle returns a copy of the string s with all Unicode letters that begin
func ToTitle(s string) string {
	return gostrings.ToTitle(s)
}
