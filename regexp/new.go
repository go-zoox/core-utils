package regexp

import "regexp"

// RegExp is the regular expression interface.
type RegExp interface {
	// Match reports whether the string s contains any match of the regular expression re.
	Match(s string) (matched bool)

	// Find is the 'All' version of Find; it returns a slice of all successive
	// matches of the expression, as defined by the 'All' description in the
	// package comment.
	// A return value of nil indicates no match.
	Find(s string) []string

	// FindN finds the matched value, max length n.
	FindN(s string, n int) []string

	// FindOne returns the first one matched value.
	FindOne(s string) string

	// Split slices s into substrings separated by the expression and returns a slice of
	// the substrings between those expression matches.
	Split(s string) []string

	// SplitN slices s into substrings separated by the expression and returns a slice of
	// the substrings between those expression matches.
	SplitN(s string, n int) []string

	// Replace returns a copy of src, replacing matches of the Regexp with the replacement text repl.
	// Inside repl, $ signs are interpreted as in Expand, so for instance $1 represents the text of the first submatch.
	Replace(src string, repl string) string

	// ReplaceFunc returns a copy of src in which all matches of the
	// Regexp have been replaced by the return value of function repl applied
	// to the matched substring. The replacement returned by repl is substituted
	// directly, without using Expand.
	ReplaceFunc(src string, repl func(string) string) string
}

type regexpX struct {
	re *regexp.Regexp
}

// New creates a new RegExp.
func New(pattern string) (RegExp, error) {
	re, err := regexp.Compile(pattern)
	if err != nil {
		return nil, err
	}

	return &regexpX{
		re: re,
	}, nil
}

func (r *regexpX) Match(s string) bool {
	return r.re.Match([]byte(s))
}

func (r *regexpX) Find(s string) []string {
	return r.FindN(s, -1)
}

func (r *regexpX) FindN(s string, n int) []string {
	return r.re.FindAllString(s, n)
}

func (r *regexpX) FindOne(s string) string {
	return r.re.FindString(s)
}

func (r *regexpX) Split(s string) []string {
	return r.SplitN(s, -1)
}

func (r *regexpX) SplitN(s string, n int) []string {
	return r.re.Split(s, n)
}

func (r *regexpX) Replace(src string, repl string) string {
	return r.re.ReplaceAllString(src, repl)
}

func (r *regexpX) ReplaceFunc(src string, repl func(string) string) string {
	return r.re.ReplaceAllStringFunc(src, repl)
}
