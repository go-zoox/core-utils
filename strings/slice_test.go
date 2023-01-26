package strings

import "testing"

func TestSlice(t *testing.T) {
	if v := Slice("abcdef", 1, 3); v != "bc" {
		t.Errorf("Expected bc, got %v", v)
	}

	if v := Slice("abcdef", 1, -1); v != "bcde" {
		t.Errorf("Expected bcdef, got %v", v)
	}

	if v := Slice("abcdef", 1, 10); v != "bcdef" {
		t.Errorf("Expected bcdef, got %v", v)
	}

	if v := Slice("abcdef", 1, 0); v != "" {
		t.Errorf("Expected \"\", got %v", v)
	}

	if v := Slice("abcdef", 0, 0); v != "" {
		t.Errorf("Expected \"\", got %v", v)
	}

	if v := Slice("abcdef", 0, -1); v != "abcde" {
		t.Errorf("Expected abcde, got %v", v)
	}

	if v := Slice("abcdef", -1, -1); v != "abcde" {
		t.Errorf("Expected \"\", got %v", v)
	}

	if v := Slice("abcdef", -1, 10); v != "abcdef" {
		t.Errorf("Expected abcdef, got %v", v)
	}

	if v := Slice("abcdef", -1, 1); v != "a" {
		t.Errorf("Expected a, got %v", v)
	}
}
