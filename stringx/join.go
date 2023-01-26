package strings

import gostrings "strings"

// Join joins the elements of a to create a single string. The separator string
// sep is placed between elements in the resulting string.
func Join(a []string, sep string) string {
	return gostrings.Join(a, sep)
}
