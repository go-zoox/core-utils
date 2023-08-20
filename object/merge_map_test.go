package object

import "testing"

func TestMergeMap(t *testing.T) {
	maps := []map[string]int{
		{
			"a": 1,
			"b": 2,
		},
		{
			"c": 3,
		},
	}

	result := MergeMap(maps...)

	if len(result) != 3 {
		t.Errorf("Expected %d values, got %d", 3, len(result))
	}

	for key, value := range result {
		if key == "a" && value != 1 {
			t.Errorf("Expected %d, got %d", 1, value)
		}

		if key == "b" && value != 2 {
			t.Errorf("Expected %d, got %d", 2, value)
		}

		if key == "c" && value != 3 {
			t.Errorf("Expected %d, got %d", 3, value)
		}
	}
}
