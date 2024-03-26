package object

import (
	"reflect"
	"testing"
)

func TestStructToMap(t *testing.T) {
	type User struct {
		ID   int
		Name string
	}

	user := User{
		ID:   1,
		Name: "John Doe",
	}

	mapValue, err := StructToMap[string, any](user)
	if err != nil {
		t.Fatal(err)
	}

	want := map[string]any{
		"ID":   float64(1),
		"Name": "John Doe",
	}

	if !reflect.DeepEqual(mapValue, want) {
		t.Fatalf("got: %v, want: %v", mapValue, want)
	}
}
