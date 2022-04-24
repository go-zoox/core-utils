package object

import "testing"

func TestOmit(t *testing.T) {
	object := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}

	values := Omit(object, "a", "b")

	if len(values) != 1 {
		t.Errorf("Expected %d values, got %d", 1, len(values))
	}

	for _, value := range values {
		if value != 3 {
			t.Errorf("Expected %d, got %d", 3, value)
		}
	}
}

func TestOmitBy(t *testing.T) {
	object := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}

	values := OmitBy(object, func(value int, key string) bool {
		return value == 1
	})

	if len(values) != 2 {
		t.Errorf("Expected %d values, got %d", 2, len(values))
	}

	for _, value := range values {
		if value != 2 && value != 3 {
			t.Errorf("Expected %d, got %d", 2, value)
		}
	}
}
