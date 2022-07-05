package safe

import "testing"

func TestSafeList(t *testing.T) {
	list := NewList()
	list.Push(1)
	list.Push(2)
	list.Push(3)
	list.Push(4)

	if v := list.Pop(); v != 4 {
		t.Errorf("Expected 4, got %v", v)
	}

	if v := list.Pop(); v != 3 {
		t.Errorf("Expected 3, got %v", v)
	}

	if v := list.First(); v != 1 {
		t.Errorf("Expected 1, got %v", v)
	}

	if v := list.Last(); v != 2 {
		t.Errorf("Expected 2, got %v", v)
	}

	if list.Size() != 2 {
		t.Errorf("Expected 2, got %v", list.Size())
	}

	if list.Length() != 2 {
		t.Errorf("Expected 2, got %v", list.Length())
	}

	list.Clear()
	if list.Size() != 0 {
		t.Errorf("Expected 0, got %v", list.Size())
	}

	if list.Length() != 0 {
		t.Errorf("Expected 0, got %v", list.Length())
	}

	list.Push(1)
	list.Push(2)
	list.Unshift(3)
	list.Unshift(4)
	if list.Size() != 4 {
		t.Errorf("Expected 4, got %v", list.Size())
	}

	if list.Length() != 4 {
		t.Errorf("Expected 4, got %v", list.Length())
	}

	if v := list.Pop(); v != 2 {
		t.Errorf("Expected 2, got %v", v)
	}

	if v := list.First(); v != 4 {
		t.Errorf("Expected 4, got %v", v)
	}

	x := list.Reduce(func(all interface{}, v interface{}) interface{} {
		return all.(int) + v.(int)
	})
	if x.(int) != 8 {
		t.Errorf("Expected 8, got %v", x)
	}
}
