package object

// Pick returns a new map with keys picked from the map.
func Pick[K comparable, V any](object map[K]V, keys ...K) map[K]V {
	result := make(map[K]V)

	for _, key := range keys {
		if value, ok := object[key]; ok {
			result[key] = value
		}
	}

	return result
}

// PickBy returns a new map by picked function from the map.
func PickBy[K comparable, V any](object map[K]V, fn func(V, K) bool) map[K]V {
	result := make(map[K]V)

	for key, value := range object {
		if fn(value, key) {
			result[key] = value
		}
	}

	return result
}
