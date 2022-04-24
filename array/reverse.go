package array

// Reverse returns an array of reversed values.
func Reverse[T any](collection []T) []T {
	ncollection := make([]T, len(collection))

	for i, j := 0, len(collection)-1; i < j; i, j = i+1, j-1 {
		ncollection[i], ncollection[j] = collection[j], collection[i]
	}

	return ncollection
}
