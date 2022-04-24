package array

import "testing"

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
