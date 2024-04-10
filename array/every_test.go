package array

import "testing"

func TestEvery(t *testing.T) {
	a := []int{4, 3, 8, 9, 10, 1, 2, 6, 7, 5}

	ok := Every(a, func(element int, index int) bool {
		return element < 11
	})
	if !ok {
		t.Error("Expected true, got", ok)
	}

	ok = Every(a, func(element int, index int) bool {
		return element < 10
	})
	if ok {
		t.Error("Expected false, got", ok)
	}
}
