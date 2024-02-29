package array

import "testing"

func TestFindIndex(t *testing.T) {
	collection := []int{1, 2, 3, 4, 5}
	// value := 3

	if index := FindIndex(collection, func(value int, index int) bool {
		return value == 3
	}); index != 2 {
		t.Error("Expected 2, got", index)
	}

	if index := FindIndex(collection, func(value int, index int) bool {
		return value == 6
	}); index != -1 {
		t.Error("Expected -1, got", index)
	}
}
