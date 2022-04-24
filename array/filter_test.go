package array

import (
	"testing"
)

func TestFilter(t *testing.T) {
	values := []int{1, 2, 3, 4, 5}
	nvalues := Filter(values, func(value int, index int) bool {
		return value > 2
	})

	if len(nvalues) != 3 {
		t.Errorf("Expected %d values, got %d", 3, len(nvalues))
	}

	for index, value := range nvalues {
		if value != values[index+2] {
			t.Errorf("Expected %d, got %d", values[index+2], value)
		}
	}
}
