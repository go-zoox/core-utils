package array

// ForEach iterates the value of slice.
func ForEach[K any](values []K, fn func(value K, index int)) {
	for index, value := range values {
		fn(value, index)
	}
}
