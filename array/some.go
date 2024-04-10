package array

// Some returns true if some of the elements is matched
func Some[T any](elements []T, compare func(element T, index int) bool) (ok bool) {
	for index, element := range elements {
		if compare(element, index) {
			return true
		}
	}

	return false
}
