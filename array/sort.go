package array

import (
	"sort"
)

// Sort sorts the generic slice.
//
//	immutable
func Sort[T any](collection []T, cmp func(v1 T, v2 T) bool) []T {
	newCollection := make([]T, len(collection))
	copy(newCollection, collection)

	sort.Slice(newCollection, func(i, j int) bool {
		return cmp(newCollection[i], newCollection[j])
	})

	return newCollection
}
