package strings

import "testing"

func TestEquals(t *testing.T) {
	if !Equals("a", "a") {
		t.Errorf("Expected true, got false")
	}

	if Equals("a", "b") {
		t.Errorf("Expected false, got true")
	}
}
