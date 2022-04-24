package array

// Max returns the maxium value in collection.
func Max[T Number](collection []T) T {
	max := collection[0]

	for _, item := range collection {
		if max < item {
			max = item
		}
	}

	return max
}
