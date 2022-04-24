package array

// FindIndex returns the index of the first occurrence of value in collection.
// If value is not present in collection, then FindIndex returns -1.
func FindIndex[T comparable](collection []T, value T) int {
	return IndexOf(collection, value)
}
