package array

import "testing"

func TestLast(t *testing.T) {
	collection := []int{1, 2, 3, 4, 5}
	if last := Last(collection); last != 5 {
		t.Error("Expected 5, got", last)
	}
}
