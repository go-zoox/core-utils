package object

import (
	"testing"
)

func TestGet(t *testing.T) {
	object := map[string]any{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": map[string]any{
			"e": map[string]any{
				"f": 4,
			},
			"f": 5,
		},
	}

	if Get(object, "a") != 1 {
		t.Errorf("Expected %d, got %d", 1, Get(object, "a"))
	}

	// if Get(object, "d") != 0 {
	// 	t.Errorf("Expected %d, got %d", 0, Get(object, "d"))
	// }

	if Get(object, "d.e.f") != 4 {
		t.Errorf("Expected %d, got %d", 4, Get(object, "d.e.f").(int))
	}

	if Get(object, "d.f") != 5 {
		t.Errorf("Expected %d, got %d", 5, Get(object, "d.f"))
	}
}
