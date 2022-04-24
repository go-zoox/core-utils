package array

import "testing"

func TestIndexOf(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	if IndexOf(a, 5) != 4 {
		t.Error("Expected 4, got", IndexOf(a, 5))
	}

	if IndexOf(a, 11) != -1 {
		t.Error("Expected -1, got", IndexOf(a, 11))
	}

	if IndexOf(a, 0) != -1 {
		t.Error("Expected -1, got", IndexOf(a, 0))
	}

	if IndexOf(a, -1) != -1 {
		t.Error("Expected -1, got", IndexOf(a, -1))
	}

	if IndexOf(a, 1) != 0 {
		t.Error("Expected 0, got", IndexOf(a, 1))
	}
}
