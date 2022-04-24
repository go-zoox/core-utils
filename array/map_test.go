package array

import (
	"testing"
)

func TestMap(t *testing.T) {
	values := []int{1, 2, 3, 4, 5}
	nvalues := Map(values, func(value int, index int) int {
		return value * 2
	})

	if len(nvalues) != len(values) {
		t.Errorf("Expected %d values, got %d", len(values), len(nvalues))
	}

	for index, value := range nvalues {
		if value != values[index]*2 {
			t.Errorf("Expected %d, got %d", values[index]*2, value)
		}
	}
}
