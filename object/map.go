package object

// Map returns a new map with the same keys by iteratee function from the map.
func Map[K comparable, V any](object map[K]V, iteratee func(V, K) V) map[K]V {
	result := make(map[K]V)

	for key, value := range object {
		result[key] = iteratee(value, key)
	}

	return result
}
