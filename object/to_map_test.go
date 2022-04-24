package object

import (
	"testing"
)

func TestToMap(t *testing.T) {
	type User struct {
		Name string
		Age  int
	}

	user := &User{
		Name: "Zero",
		Age:  18,
	}

	userMap, err := ToMap(user)
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}

	if userMap["Name"] != "Zero" {
		t.Errorf("Expected %s, got %s", "Zero", userMap["Name"])
	}

	if userMap["Age"] != 18 {
		t.Errorf("Expected %d, got %d", 18, userMap["Age"])
	}
}

func TestToMapWithTag(t *testing.T) {
	type User struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	user := &User{
		Name: "Zero",
		Age:  18,
	}

	userMap, err := ToMap(user, "json")
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}

	if userMap["name"] != "Zero" {
		t.Errorf("Expected %s, got %s", "Zero", userMap["name"])
	}

	if userMap["age"] != 18 {
		t.Errorf("Expected %d, got %d", 18, userMap["age"])
	}
}

func TestToMapNestStruct(t *testing.T) {
	type User struct {
		Name    string `json:"name"`
		Age     int    `json:"age"`
		Address struct {
			City string `json:"city"`
		} `json:"address"`
	}

	user := &User{
		Name: "Zero",
		Age:  18,
		Address: struct {
			City string `json:"city"`
		}{
			City: "New York",
		},
	}

	userMap, err := ToMap(user, "json")
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}

	if userMap["name"] != "Zero" {
		t.Errorf("Expected %s, got %s", "Zero", userMap["name"])
	}

	if userMap["age"] != 18 {
		t.Errorf("Expected %d, got %d", 18, userMap["age"])
	}

	// fmt.PrintJSON(userMap)

	if Get(userMap, "address.city") != "New York" {
		t.Errorf("Expected %s, got %s", "New York", Get(userMap, "address.city"))
	}
}

func TestToMapNestSlice(t *testing.T) {
	type User struct {
		Name    string   `json:"name"`
		Age     int      `json:"age"`
		Hobbies []string `json:"hobbies"`
		Address struct {
			City string `json:"city"`
		} `json:"address"`
	}

	user := &User{
		Name: "Zero",
		Age:  18,
		Hobbies: []string{
			"reading",
			"coding",
		},
		Address: struct {
			City string `json:"city"`
		}{
			City: "New York",
		},
	}

	userMap, err := ToMap(user, "json")
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}

	// fmt.PrintJSON(userMap)

	// gofmt.Println(Get(userMap, "hobbies.0"))
	// gofmt.Println(Get(userMap, "hobbies.1"))
	// gofmt.Println(Get(userMap, "hobbies.2"))
	// gofmt.Println(Get(userMap, "address.city"))

	if Get(userMap, "hobbies.0") != "reading" {
		t.Errorf("Expected %s, got %s", "reading", Get(userMap, "hobbies[0]"))
	}

	if Get(userMap, "hobbies.1") != "coding" {
		t.Errorf("Expected %s, got %s", "coding", Get(userMap, "hobbies[1]"))
	}

	if Get(userMap, "hobbies.2") != nil {
		t.Errorf("Expected nil, got %s", Get(userMap, "hobbies[2]"))
	}

	if Get(userMap, "address.city") != "New York" {
		t.Errorf("Expected %s, got %s", "New York", Get(userMap, "address.city"))
	}
}
