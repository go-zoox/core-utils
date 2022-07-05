package stringx

import (
	"math/rand"
	"time"
)

// random returns a random string of length n.
func Random(n int, seeds ...string) string {
	rand.Seed(time.Now().UnixNano())

	var seed string
	if len(seeds) > 0 {
		seed = seeds[0]
	} else {
		seed = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	}

	var bytes = make([]byte, n)
	for i := range bytes {
		bytes[i] = seed[rand.Intn(len(seed))]
	}
	return string(bytes)
}
