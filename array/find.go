package array

// Find searches an element in a slice based on a predicated.
// It returns element and true if element was found.
func Find[T any](elements []T, compare func(element T, index int) bool) (T, bool) {
	for index, element := range elements {
		if compare(element, index) {
			return element, true
		}
	}

	var empty T
	return empty, false
}
