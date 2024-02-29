package array

// FindIndex returns the index of the first occurrence of value in collection.
// If value is not present in collection, then FindIndex returns -1.
func FindIndex[T any](elements []T, compare func(element T, index int) bool) int {
	for index, element := range elements {
		if compare(element, index) {
			return index
		}
	}

	return -1
}
