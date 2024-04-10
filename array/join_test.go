package array

import (
	"testing"

	"github.com/go-zoox/core-utils/cast"
	"github.com/go-zoox/testify"
)

func TestJoinWithString(t *testing.T) {
	values := []string{"a", "b", "c", "d", "e"}
	separator := ","
	expected := "a,b,c,d,e"

	received, err := Join(values, separator)
	// testify.Nil(t, err)
	if err != nil {
		t.Fatalf("Error: %s", err)
	}

	testify.Equal(t, expected, received)
}

func TestJoinWithInt(t *testing.T) {
	values := []int{1, 2, 3, 4, 5}
	separator := ","
	expected := "1,2,3,4,5"

	received, err := Join(values, separator)
	// testify.Nil(t, err)
	if err != nil {
		t.Fatalf("Error: %s", err)
	}

	testify.Equal(t, expected, received)
}

func TestJoinWithIntAndFunction(t *testing.T) {
	values := []int{1, 2, 3, 4, 5}
	separator := ","
	expected := "1,2,3,4,5"

	received, err := Join(values, separator, func(value int, index int) string {
		return cast.ToString(value)
	})
	// testify.Nil(t, err)
	if err != nil {
		t.Fatalf("Error: %s", err)
	}

	testify.Equal(t, expected, received)
}

func TestJoinWithStructAndFunction(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}

	values := []Person{
		{"Alice", 25},
		{"Bob", 30},
		{"Charlie", 35},
	}
	separator := ","
	expected := "Alice,Bob,Charlie"

	received, err := Join(values, separator, func(value Person, index int) string {
		return value.Name
	})
	// testify.Nil(t, err)
	if err != nil {
		t.Fatalf("Error: %s", err)
	}

	testify.Equal(t, expected, received)
}

func TestJoinWithStructAndFunctionError(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}

	values := []Person{
		{"Alice", 25},
		{"Bob", 30},
		{"Charlie", 35},
	}
	separator := ","
	expected := "Alice,Bob,Charlie"

	received, err := Join(values, separator, func(value Person, index int) string {
		return value.Name
	})
	// testify.Nil(t, err)
	if err != nil {
		t.Fatalf("Error: %s", err)
	}

	testify.Equal(t, expected, received)
}
