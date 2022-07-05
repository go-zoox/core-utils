package stringx

import "testing"

func TestPadZero(t *testing.T) {
	s := "123"
	if PadZero(s, 5) != "12300" {
		t.Error("PadZero failed")
	}
}

func TestPadLeft(t *testing.T) {
	s := "123"
	if PadLeft(s, 5, "#") != "##123" {
		t.Error("PadLeft failed")
	}
}

func TestPadRight(t *testing.T) {
	s := "123"
	if PadRight(s, 5, "#") != "123##" {
		t.Error("PadRight failed")
	}
}
