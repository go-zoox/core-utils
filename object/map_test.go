package object

import "testing"

func TestMap(t *testing.T) {
	object := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}

	values := Map(object, func(value int, key string) int {
		return value * 2
	})

	if len(values) != len(object) {
		t.Errorf("Expected %d values, got %d", len(object), len(values))
	}

	for key, value := range values {
		if value != object[key]*2 {
			t.Errorf("Expected %d, got %d", object[key]*2, value)
		}
	}
}
