package array

// Includes returns true if the array contains the value.
func Includes[K comparable](values []K, value K) bool {
	for _, v := range values {
		if v == value {
			return true
		}
	}

	return false
}
