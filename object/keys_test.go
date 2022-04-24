package object

import "testing"

func TestKeys(t *testing.T) {
	object := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}

	keys := Keys(object)

	if len(keys) != 3 {
		t.Errorf("Expected %d keys, got %d", 3, len(keys))
	}

	for _, key := range keys {
		if key != "a" && key != "b" && key != "c" {
			t.Errorf("Expected %s, got %s", "a", key)
		}
	}
}
