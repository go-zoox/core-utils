package object

import (
	"testing"
)

func TestGet(t *testing.T) {
	object := map[string]any{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": map[string]any{
			"e": map[string]any{
				"f": 4,
			},
			"f": 5,
		},
	}

	if Get(object, "a") != 1 {
		t.Errorf("Expected %d, got %d", 1, Get(object, "a"))
	}

	// if Get(object, "d") != 0 {
	// 	t.Errorf("Expected %d, got %d", 0, Get(object, "d"))
	// }

	if Get(object, "d.e.f") != 4 {
		t.Errorf("Expected %d, got %d", 4, Get(object, "d.e.f").(int))
	}

	if Get(object, "d.f") != 5 {
		t.Errorf("Expected %d, got %d", 5, Get(object, "d.f"))
	}
}

func TestGet2(t *testing.T) {
	object := map[string]interface{}{
		"app_name":  "gozoox",
		"log_level": "DEBUG",
		"redis": map[string]interface{}{
			"ip":   "127.0.0.1",
			"port": "6739",
			// "port": 6739,
		},
		"ports": []int64{
			6739,
			6740,
		},
		"maps": map[string]interface{}{
			"key1": "value1",
			"key2": "value2",
		},
		"providers": map[string]any{
			"google": map[string]any{
				"client_id":     "google_client_id",
				"client_secret": "google_client_secret",
			},
			"facebook": map[string]any{
				"client_id":     "facebook_client_id",
				"client_secret": "facebook_client_secret",
			},
			"github": map[string]any{
				"client_id":     "github_client_id",
				"client_secret": "github_client_secret",
			},
		},
		"users": []map[string]any{
			{
				"name": "user1",
				"age":  18,
			},
			{
				"name": "user2",
				"age":  20,
			},
		},
		"type_transform": "666",
	}

	if Get(object, "providers.google.client_id") != "google_client_id" {
		t.Errorf("Expected %s, got %s", "google_client_id", Get(object, "providers.google.client_id").(string))
	}
}

func TestGet3(t *testing.T) {
	object := map[string]interface{}{
		"app_name":  "gozoox",
		"log_level": "DEBUG",
		"redis": map[string]interface{}{
			"ip":   "127.0.0.1",
			"port": "6739",
			// "port": 6739,
		},
		"ports": []int64{
			6739,
			6740,
		},
		"maps": map[string]interface{}{
			"key1": "value1",
			"key2": "value2",
		},
		"providers": map[string]any{
			"google": map[string]any{
				"client_id":     "google_client_id",
				"client_secret": "google_client_secret",
			},
			"facebook": map[string]any{
				"client_id":     "facebook_client_id",
				"client_secret": "facebook_client_secret",
			},
			"github": map[string]any{
				"client_id":     "github_client_id",
				"client_secret": "github_client_secret",
			},
		},
		"users": []map[string]any{
			{
				"name": "user1",
				"age":  18,
			},
			{
				"name": "user2",
				"age":  20,
			},
		},
		"type_transform": "666",
	}

	if Get(object, "ports.0").(int64) != 6739 {
		t.Errorf("Expected %d, got %d", 6739, Get(object, "ports.0"))
	}

	if Get(object, "ports.1").(int64) != 6740 {
		t.Errorf("Expected %d, got %d", 6740, Get(object, "ports.1"))
	}

	if Get(object, "users.0.name") != "user1" {
		t.Errorf("Expected %s, got %s", "user1", Get(object, "users.0.name").(string))
	}

	if Get(object, "users.1.name") != "user2" {
		t.Errorf("Expected %s, got %s", "user2", Get(object, "users.1.name").(string))
	}
}

// reflect: call of reflect.Value.Interface on zero Value
func TestGet4(t *testing.T) {
	value := map[string]interface{}{
		"foo8": nil,
	}

	// testify.Equal(t, false, Get(value, "foo8").(bool))
	// testify.Equal(t, false, Get(value, "foo9").(bool))

	if Get(value, "foo8") != nil {
		t.Errorf("Expected %v, got %s", nil, Get(value, "foo8"))
	}

	if Get(value, "foo9") != nil {
		t.Errorf("Expected %v, got %s", nil, Get(value, "foo9"))
	}
}
