package array

import "testing"

func TestSome(t *testing.T) {
	a := []int{4, 3, 8, 9, 10, 1, 2, 6, 7, 5}

	ok := Some(a, func(element int, index int) bool {
		return element == 10
	})
	if !ok {
		t.Error("Expected true, got", ok)
	}

	ok = Some(a, func(element int, index int) bool {
		return element == 11
	})
	if ok {
		t.Error("Expected false, got", ok)
	}
}
