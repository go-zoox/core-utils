package regexp

import "regexp"

// Match reports whether the string s
// contains any match of the regular expression pattern.
// More complicated queries need to use Compile and the full Regexp interface.
func Match(pattern string, s string) (matched bool) {
	ok, err := regexp.MatchString(pattern, s)

	if err != nil {
		return false
	}

	return ok
}
