package stringx

import (
	"fmt"
	"testing"
)

func TestRandom(t *testing.T) {
	if v := Random(10); len(v) != 10 {
		t.Errorf("Expected 10, got %v", len(v))
	} else {
		fmt.Println(v)
	}

	if v := Random(10, "abc"); len(v) != 10 {
		t.Errorf("Expected 10, got %v", len(v))
	} else {
		fmt.Println(v)
	}

	if v := Random(10, "abcdef"); len(v) != 10 {
		t.Errorf("Expected 10, got %v", len(v))
	} else {
		fmt.Println(v)
	}
}
