package array

// IndexOf returns the index of the first occurrence of value in collection.
// If value is not present in collection, then IndexOf returns -1.
func IndexOf[T comparable](collection []T, value T) int {
	for index, item := range collection {
		if item == value {
			return index
		}
	}

	return -1
}
