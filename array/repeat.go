package array

// Repeat creates a slice with N copies of the initail value.
func Repeat[T any](count int, value T) []T {
	result := make([]T, 0, count)

	for i := 0; i < count; i++ {
		result = append(result, value)
	}

	return result
}
