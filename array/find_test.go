package array

import "testing"

func TestFind(t *testing.T) {
	values := []int{1, 2, 3, 4, 5}
	result, found := Find(values, func(value int, index int) bool {
		return value == 3
	})

	if !found {
		t.Errorf("Expected %d to be found", 3)
	}

	if result != 3 {
		t.Errorf("Expected %d, got %d", 3, result)
	}
}
