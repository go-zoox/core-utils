package array

// LastIndexOf returns the index of the last occurrence of value in collection.
// If value is not present in collection, then LastIndexOf returns -1.
func LastIndexOf[T comparable](collection []T, value T) int {
	for index := len(collection) - 1; index >= 0; index-- {
		if collection[index] == value {
			return index
		}
	}

	return -1
}
