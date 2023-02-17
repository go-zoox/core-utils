package array

// Reverse returns an array of reversed values.
func Reverse[T any](collection []T) []T {
	length := len(collection)
	half := length/2 + 1
	ncollection := make([]T, length)

	for i := 0; i < half; i++ {
		j := length - i - 1
		ncollection[i], ncollection[j] = collection[j], collection[i]
	}

	return ncollection
}
