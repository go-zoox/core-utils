package array

import "testing"

func TestGroupBy(t *testing.T) {
	values := []int{1, 2, 3, 4, 5}

	grouped := GroupBy(values, func(value int, index int) string {
		mod := value % 2
		if mod == 0 {
			return "even"
		}

		return "odd"
	})

	if groupEven, ok := grouped["even"]; !ok {
		t.Errorf("Expected group 'even'")
	} else {
		if len(groupEven) != 2 {
			t.Errorf("Expected %d values, got %d", 2, len(groupEven))
		}

		for _, value := range groupEven {
			if value != 2 && value != 4 {
				t.Errorf("Expected 2,4, got %d", value)
			}
		}
	}
}
