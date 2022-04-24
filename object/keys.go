package object

// Keys returns the keys of the map.
func Keys[K comparable, V any](object map[K]V) []K {
	keys := make([]K, 0)

	for key := range object {
		keys = append(keys, key)
	}

	return keys
}
