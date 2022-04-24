package object

import "testing"

func TestWalk(t *testing.T) {
	user := map[string]any{
		"name": "John",
		"age":  30,
		"address": map[string]any{
			"city": "New York",
			"street": map[string]any{
				"name":   "Wall Street",
				"number": 12,
			},
		},
		"houses": []any{
			map[string]any{
				"name": "Buckingham Palace",
			},
			map[string]any{
				"name": "White House",
			},
		},
	}

	indexes := map[string]any{}

	Walk(user, func(v any, k string, path string) {
		indexes[path] = v
	})

	if indexes["name"] != "John" {
		t.Error("Expected John, got", indexes["name"])
	}

	if indexes["age"] != 30 {
		t.Error("Expected 30, got", indexes["age"])
	}

	if indexes["address.city"] != "New York" {
		t.Error("Expected New York, got", indexes["address.city"])
	}

	if indexes["address.street.name"] != "Wall Street" {
		t.Error("Expected Wall Street, got", indexes["address.street.name"])
	}

	if indexes["address.street.number"] != 12 {
		t.Error("Expected 12, got", indexes["address.street.number"])
	}

	if indexes["houses.0.name"] != "Buckingham Palace" {
		t.Error("Expected Buckingham Palace, got", indexes["houses.0.name"])
	}

	if indexes["houses.1.name"] != "White House" {
		t.Error("Expected White House, got", indexes["houses.1.name"])
	}
}
