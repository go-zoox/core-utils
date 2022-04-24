package array

import (
	"testing"
)

func TestIncludes(t *testing.T) {
	values := []int{1, 2, 3, 4, 5}

	if !Includes(values, 3) {
		t.Errorf("Expected %d to be found", 3)
	}

	if Includes(values, 6) {
		t.Errorf("Expected %d to not be found", 6)
	}
}
