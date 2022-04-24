package array

import "testing"

func TestLastIndexOf(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	if LastIndexOf(a, 5) != 4 {
		t.Error("Expected 4, got", LastIndexOf(a, 5))
	}

	if LastIndexOf(a, 11) != -1 {
		t.Error("Expected -1, got", LastIndexOf(a, 11))
	}

	if LastIndexOf(a, 0) != -1 {
		t.Error("Expected -1, got", LastIndexOf(a, 0))
	}

	if LastIndexOf(a, -1) != -1 {
		t.Error("Expected -1, got", LastIndexOf(a, -1))
	}

	if LastIndexOf(a, 1) != 0 {
		t.Error("Expected 0, got", LastIndexOf(a, 1))
	}
}
