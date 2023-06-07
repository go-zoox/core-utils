package array

import (
	"testing"
)

func TestReverse(t *testing.T) {
	origin := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	reversed := Reverse(origin)

	if len(reversed) != 10 {
		t.Error("Expected 10, got", len(reversed))
	}

	for i, v := range reversed {
		if v != 10-i {
			t.Error("Expected", 10-i, "got", v)
		}
	}
}

func TestReverse2(t *testing.T) {
	origin := []string{"1", "2", "3"}
	reversed := Reverse(origin)

	if len(reversed) != 3 {
		t.Error("Expected 3, got", len(reversed))
	}

	if reversed[0] != "3" {
		t.Error("Expected 3, got", reversed[0])
	}
	if reversed[1] != "2" {
		t.Error("Expected 2, got", reversed[1])
	}
	if reversed[2] != "1" {
		t.Error("Expected 1, got", reversed[2])
	}
}
