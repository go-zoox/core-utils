package array

import "testing"

func TestMin(t *testing.T) {
	collection := []int{1, 2, 3, 4, 5}
	if min := Min(collection); min != 1 {
		t.Error("Expected 1, got", min)
	}
}
