package object

import (
	"testing"

	"github.com/go-zoox/testify"
)

func TestUnique(t *testing.T) {
	source := []string{"a", "b", "c", "a", "b"}

	target := Unique(source)

	if len(target) != 3 {
		t.Errorf("Expected %d values, got %d", 3, len(target))
	}

	testify.Equal(t, "a", target[0])
	testify.Equal(t, "b", target[1])
	testify.Equal(t, "c", target[2])
}

func TestUniqueBy(t *testing.T) {
	type User struct {
		Name string
		Age  int64
	}

	source := []User{
		{
			Name: "Zero",
			Age:  18,
		},
		{
			Name: "Amy",
			Age:  16,
		},
		{
			Name: "Zero",
			Age:  100,
		},
	}

	target := UniqueBy(source, func(origin User, i int) string {
		return origin.Name
	})
	if len(target) != 2 {
		t.Errorf("Expected %d values, got %d", 3, len(target))
	}

	testify.Equal(t, "Zero", target[0].Name)
	testify.Equal(t, 18, target[0].Age)

	testify.Equal(t, "Amy", target[1].Name)
	testify.Equal(t, 16, target[1].Age)
}
