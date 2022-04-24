package array

// Last returns the last element of collection.
// If collection is empty, then Last returns nil.
func Last[T any](collection []T) T {
	if len(collection) == 0 {
		var empty T
		return empty
	}

	return collection[len(collection)-1]
}
