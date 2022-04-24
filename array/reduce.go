package array

// Reduce returns the result of applying the function fn to each value in the array and an accumulator, starting with the initial value.
func Reduce[K any, R any](values []K, fn func(all R, value K, index int) R, initialValue R) R {
	result := initialValue

	if len(values) == 0 {
		return result
	}

	for index, value := range values {
		result = fn(result, value, index)
	}

	return result
}
