package array

import (
	"math/rand"
	"time"
)

// Shuffle returns an array of shuffled values. Use the Fisher-Yates shuffle algorithm.
func Shuffle[T any](collection []T) []T {
	ncollection := make([]T, len(collection))

	// set a new seed
	rand.Seed(int64(time.Now().Nanosecond()))

	rand.Shuffle(len(collection), func(i, j int) {
		ncollection[i], ncollection[j] = collection[j], collection[i]
	})

	return ncollection
}
