package safe

import "testing"

func TestSafeList(t *testing.T) {
	list := NewList[int]()
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

	x := list.Reduce(func(all int, v int) int {
		return all + v
	})
	if x != 8 {
		t.Errorf("Expected 8, got %v", x)
	}
}

func TestListCapacity(t *testing.T) {
	capacity := 3
	l := NewList[int](func(lc *ListConfig) {
		lc.Capacity = capacity
	})
	l.Push(1)
	l.Push(2)
	l.Push(3)
	if l.Size() != 3 {
		t.Fatalf("expected 3, got %v", l.Size())
	}

	l.Push(4)
	if l.Size() != 3 {
		t.Fatalf("expected 3, got %v", l.Size())
	}
	if l.data[0] != 2 {
		t.Fatalf("expected 2, got %v", l.data[0])
	}
	if l.data[1] != 3 {
		t.Fatalf("expected 3, got %v", l.data[0])
	}
	if l.data[2] != 4 {
		t.Fatalf("expected 4, got %v", l.data[0])
	}

	l.Push(5)
	if l.Size() != 3 {
		t.Fatalf("expected 3, got %v", l.Size())
	}
	if l.data[0] != 3 {
		t.Fatalf("expected 3, got %v", l.data[0])
	}
	if l.data[1] != 4 {
		t.Fatalf("expected 4, got %v", l.data[0])
	}
	if l.data[2] != 5 {
		t.Fatalf("expected 5, got %v", l.data[0])
	}

	l.Unshift(6)
	if l.Size() != 3 {
		t.Fatalf("expected 3, got %v", l.Size())
	}
	if l.data[0] != 6 {
		t.Fatalf("expected 6, got %v", l.data[0])
	}
	if l.data[1] != 3 {
		t.Fatalf("expected 3, got %v", l.data[0])
	}
	if l.data[2] != 4 {
		t.Fatalf("expected 4, got %v", l.data[0])
	}

	l.Pop()
	if l.Size() != 2 {
		t.Fatalf("expected 2, got %v", l.Size())
	}
	if l.data[0] != 6 {
		t.Fatalf("expected 6, got %v", l.data[0])
	}
	if l.data[1] != 3 {
		t.Fatalf("expected 3, got %v", l.data[0])
	}

	l.Shift()
	if l.Size() != 1 {
		t.Fatalf("expected 1, got %v", l.Size())
	}
	if l.data[0] != 3 {
		t.Fatalf("expected 3, got %v", l.data[0])
	}
}

func TestListWithCustomTypeComparable(t *testing.T) {
	type Obj struct {
		Attr string
		A    interface{}
	}

	// func (o *Obj) IsEqual(other *Obj) int {
	// 	return 0
	// }

	list := NewList[Obj]()

	v := Obj{Attr: "123"}
	v1 := Obj{Attr: "456"}
	list.Push(v)
	list.Push(v1)

	if list.Size() != 2 {
		t.Errorf("Expected 1, got %v", list.Size())
	}

	if list.IndexOf(v) != 0 {
		t.Errorf("Expected 0, got %v", list.IndexOf(v))
	}

	if list.IndexOf(v1) != 1 {
		t.Errorf("Expected 1, got %v", list.IndexOf(v1))
	}
}
