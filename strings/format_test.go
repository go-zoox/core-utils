package strings

import (
	"fmt"
	"testing"
)

func TestFormat(t *testing.T) {
	// Simple
	if v := Format("i am {{name}}", map[string]any{
		"name": "Zero",
	}); v != "i am Zero" {
		t.Errorf("Expected i am Zero, got %v", v)
	} else {
		fmt.Println(v)
	}

	// Multiple use on same key
	if v := Format("i am {{name}}. Are you {{name}}?", map[string]any{
		"name": "Zero",
	}); v != "i am Zero. Are you Zero?" {
		t.Errorf("Expected i am Zero. Are you Zero?, got %v", v)
	} else {
		fmt.Println(v)
	}

	// Multiple keys
	if v := Format("i am {{name}}. My age is {{age}}", map[string]any{
		"name": "Zero",
		"age":  18,
	}); v != "i am Zero. My age is 18" {
		t.Errorf("Expected i am Zero. My age is 18, got %v", v)
	} else {
		fmt.Println(v)
	}

	// nil
	if v := Format("i am {{name}}. My age is {{age}}", map[string]any{
		"name": "Zero",
	}); v != "i am Zero. My age is " {
		t.Errorf("Expected i am Zero. My age is , got %v", v)
	} else {
		fmt.Println(v)
	}

	// empty
	if v := Format("i am {{name}}. My age is {{age}}", map[string]any{
		"name": "Zero",
		"age":  "",
	}); v != "i am Zero. My age is " {
		t.Errorf("Expected i am Zero. My age is , got %v", v)
	} else {
		fmt.Println(v)
	}

	// key contains white space
	if v := Format("i am {{innormal key}}.", map[string]any{
		"innormal key": "Zero 18",
		"age":          18,
	}); v != "i am Zero 18." {
		t.Errorf("Expected i am Zero 18., got %v", v)
	} else {
		fmt.Println(v)
	}

	// nested keys
	if v := Format("i am {{name}}. My age is {{age}}. My name is {{info.name}}", map[string]any{
		"name": "Zero",
		"age":  18,
		"info": map[string]any{
			"name": "Zero",
		},
	}); v != "i am Zero. My age is 18. My name is Zero" {
		t.Errorf("Expected i am Zero. My age is 18. My name is Zero, got %v", v)
	} else {
		fmt.Println(v)
	}

	// object value => equals to ""
	if v := Format("i am {{name}}. My age is {{age}}. My name is {{info}}", map[string]any{
		"name": "Zero",
		"age":  18,
		"info": map[string]any{
			"name": "Zero",
		},
	}); v != "i am Zero. My age is 18. My name is " {
		t.Errorf("Expected i am Zero. My age is 18. My name is , got %v", v)
	} else {
		fmt.Println(v)
	}
}
