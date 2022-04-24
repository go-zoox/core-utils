package array

import "testing"

func TestRepeat(t *testing.T) {
	a := Repeat(10, "a")
	if len(a) != 10 {
		t.Error("Expected 10, got", len(a))
	}

	for _, v := range a {
		if v != "a" {
			t.Error("Expected a, got", v)
		}
	}
}
