package strings

import (
	"regexp"

	"github.com/go-zoox/core-utils/cast"
	"github.com/go-zoox/core-utils/object"
)

const FormatPattern = `{{[^}]+}}`

// Format formats a string with template bracket {{key}}} and value map.
func Format(pattern string, data map[string]any) string {
	bytes := regexp.MustCompile(FormatPattern).ReplaceAllFunc([]byte(pattern), func(b []byte) []byte {
		key := string(b[2 : len(b)-2])
		value := object.Get(data, key)

		if value == nil {
			return nil
		}

		return []byte(cast.ToString(value))
	})

	return string(bytes)
}
