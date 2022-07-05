package stringx

import (
	"testing"
)

func TestSplit(t *testing.T) {
	if v := Split("a,b,c", ","); v[0] != "a" || v[1] != "b" || v[2] != "c" {
		t.Errorf("Expected [a,b,c], got %v", v)
	}

	if v := SplitN("a,b,c", ",", 2); v[0] != "a" || v[1] != "b,c" {
		t.Errorf("Expected [], got %v", v)
	}
}
