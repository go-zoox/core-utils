package strings

import (
	"regexp"

	"github.com/go-zoox/core-utils/cast"
	"github.com/go-zoox/core-utils/object"
)

// FormatRe is the regular expression for format.
const FormatRe = `\{[^\}]+\}`

// Format formats a string with template bracket {key}}} and value map.
func Format(pattern string, data map[string]any) string {
	bytes := regexp.MustCompile(FormatRe).ReplaceAllFunc([]byte(pattern), func(b []byte) []byte {
		key := string(b[1 : len(b)-1])
		value := object.Get(data, key)

		if value == nil {
			return nil
		}

		return []byte(cast.ToString(value))
	})

	return string(bytes)
}
