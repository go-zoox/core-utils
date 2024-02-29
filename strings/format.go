package strings

import (
	"regexp"
	"strings"

	"github.com/go-zoox/core-utils/cast"
	"github.com/go-zoox/core-utils/object"
)

// FormatRe is the regular expression for format.
const FormatRe = `\{\{([^\}]+)\}\}`

// Format formats a string with template bracket {key}}} and value map.
func Format(pattern string, data map[string]any) string {
	re := regexp.MustCompile(FormatRe)

	return re.ReplaceAllStringFunc(pattern, func(m string) string {
		key := strings.Trim(m, "{}")
		value := object.Get(data, key)

		if value == nil {
			return ""
		}

		return cast.ToString(value)
	})
}
