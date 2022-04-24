package array

// Map applies the function fn to each value in the array and returns a new array with the results.
func Map[K any, R any](values []K, fn func(value K, index int) R) []R {
	nvalues := make([]R, len(values))

	for index, value := range values {
		nvalues[index] = fn(value, index)
	}

	return nvalues
}
