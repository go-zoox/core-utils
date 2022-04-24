package array

import "testing"

func TestShuffle(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	b := Shuffle(a)
	t.Log(b)
}
