package array

// Filter filters the array using the function fn.
func Filter[K any](values []K, fn func(value K, index int) bool) []K {
	nvalues := make([]K, 0)

	for index, value := range values {
		if fn(value, index) {
			nvalues = append(nvalues, value)
		}
	}

	return nvalues
}
