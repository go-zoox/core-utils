package object

import "testing"

func TestPick(t *testing.T) {
	object := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}

	values := Pick(object, "a", "b")

	if len(values) != 2 {
		t.Errorf("Expected %d values, got %d", 2, len(values))
	}

	for _, value := range values {
		if value != 1 && value != 2 {
			t.Errorf("Expected %d, got %d", 1, value)
		}
	}
}

func TestPickBy(t *testing.T) {
	object := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}

	values := PickBy(object, func(value int, key string) bool {
		return value == 1
	})

	if len(values) != 1 {
		t.Errorf("Expected %d values, got %d", 1, len(values))
	}

	for _, value := range values {
		if value != 1 {
			t.Errorf("Expected %d, got %d", 1, value)
		}
	}
}
