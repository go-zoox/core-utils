package array

// Every returns true if all of the elements is matched
func Every[T any](elements []T, compare func(element T, index int) bool) (ok bool) {
	for index, element := range elements {
		if !compare(element, index) {
			return false
		}
	}

	return true
}
