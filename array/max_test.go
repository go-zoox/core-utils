package array

import "testing"

func TestMax(t *testing.T) {
	collection := []int{1, 2, 3, 4, 5}
	if max := Max(collection); max != 5 {
		t.Error("Expected 1, got", max)
	}
}
