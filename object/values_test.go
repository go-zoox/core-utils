package object

import "testing"

func TestValues(t *testing.T) {
	object := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}

	values := Values(object)

	if len(values) != 3 {
		t.Errorf("Expected %d values, got %d", 3, len(values))
	}

	for _, value := range values {
		if value != 1 && value != 2 && value != 3 {
			t.Errorf("Expected %d, got %d", 1, value)
		}
	}
}
