package array

// GroupBy ...
func GroupBy[T any, K comparable](values []T, iteratee func(value T, index int) K) map[K][]T {
	grouped := make(map[K][]T)

	for index, value := range values {
		key := iteratee(value, index)

		if _, ok := grouped[key]; !ok {
			grouped[key] = make([]T, 0)
		}

		grouped[key] = append(grouped[key], value)
	}

	return grouped
}
