package safe

import "testing"

func TestSafeMap(t *testing.T) {
	m := NewMap()
	m.Set("a", 1)
	m.Set("b", 2)
	m.Set("c", 3)

	if v := m.Get("a"); v != 1 {
		t.Errorf("Expected 1, got %v", v)
	}

	if v := m.Get("b"); v != 2 {
		t.Errorf("Expected 2, got %v", v)
	}

	if v := m.Get("c"); v != 3 {
		t.Errorf("Expected 3, got %v", v)
	}

	if v := m.Get("d"); v != nil {
		t.Errorf("Expected nil, got %v", v)
	}

	m.Del("a")
	if v := m.Get("a"); v != nil {
		t.Errorf("Expected nil, got %v", v)
	}

	if len(m.Keys()) != 2 {
		t.Errorf("Expected 2 keys, got %v", len(m.Keys()))
	}

	// race
	i := 0
	for {
		if i > 100 {
			return
		}

		i += 1
		go m.Set("a", i)
	}
}
