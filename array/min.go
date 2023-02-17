package array

// Number ...
type Number interface {
	int | int64 | float64 | float32 | uint | uint8 | uint16 | uint32 | uint64 | int8 | int16 | int32
}

// Min returns the minimum value in collection.
func Min[T Number](collection []T) T {
	min := collection[0]

	for _, item := range collection {
		if min > item {
			min = item
		}
	}

	return min
}
