package array

import (
	"testing"
)

func TestReduce(t *testing.T) {
	values := []int{1, 2, 3, 4, 5}
	result := Reduce(values, func(all int, value int, index int) int {
		return all + value
	}, 0)

	if result != 15 {
		t.Errorf("Expected %d, got %d", 15, result)
	}
}
