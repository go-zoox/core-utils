package object

// Values returns the values of the map.
func Values[K comparable, V any](object map[K]V) []V {
	values := make([]V, 0)

	for _, value := range object {
		values = append(values, value)
	}

	return values
}
