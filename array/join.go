package array

import (
	"strings"

	"github.com/go-zoox/core-utils/cast"
)

// Join joins all elements of the array into a string.
func Join[K any](values []K, seperator string, fn ...func(value K, index int) string) (string, error) {
	nvalues := make([]string, len(values))

	if len(fn) == 0 {
		for index, value := range values {
			// if reflect.TypeOf(value).Kind() != reflect.String {
			// 	return "", fmt.Errorf("value at index %d is not a string", index)
			// }

			// nvalues[index] = reflect.ValueOf(value).String()

			nvalues[index] = cast.ToString(value)
		}

		return strings.Join(nvalues, seperator), nil
	}

	for index, value := range values {
		nvalues[index] = fn[0](value, index)
	}

	return strings.Join(nvalues, seperator), nil
}
