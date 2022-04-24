package object

import "github.com/go-zoox/core-utils/array"

// Omit returns a new map with keys omitted from the map.
func Omit[K comparable, V any](object map[K]V, keys ...K) map[K]V {
	result := make(map[K]V)

	for key, value := range object {
		if !array.Includes(keys, key) {
			result[key] = value
		}
	}

	return result
}

// OmitBy returns a new map by omitted function from the map.
func OmitBy[K comparable, V any](object map[K]V, fn func(V, K) bool) map[K]V {
	result := make(map[K]V)

	for key, value := range object {
		if !fn(value, key) {
			result[key] = value
		}
	}

	return result
}
