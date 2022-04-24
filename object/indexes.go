package object

// Indexes flattens a nest map into a map with only one level.
func Indexes[K comparable, V any](object map[K]V) map[string]V {
	indexes := make(map[string]V)

	Walk(object, func(v V, k K, path string) {
		indexes[path] = v
	})

	return indexes
}
