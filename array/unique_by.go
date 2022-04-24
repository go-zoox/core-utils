package array

// UniqueBy returns an array of unique values by given filter function.
func UniqueBy[T any, U comparable](collection []T, iteratee func(value T, index int) U) []T {
	result := make([]T, 0, len(collection))
	seen := make(map[U]struct{}, len(collection))

	for index, item := range collection {
		key := iteratee(item, index)

		if _, ok := seen[key]; !ok {
			seen[key] = struct{}{}
			result = append(result, item)
		}
	}

	return result
}
