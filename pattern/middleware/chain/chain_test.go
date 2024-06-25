package chain

import (
	"reflect"
	"testing"
)

func TestChain(t *testing.T) {
	results := []string{}
	middlewares := []func(ctx string, next func()){
		func(ctx string, next func()) {
			results = append(results, "first")
			next()
			results = append(results, "first")
		},
		func(ctx string, next func()) {
			results = append(results, "second")
			next()
			results = append(results, "second")
		},
		func(ctx string, next func()) {
			results = append(results, "third")
			next()
			results = append(results, "third")
		},
	}

	Chain("start", middlewares, func(ctx string) {
		results = append(results, "final")
	})

	expected := []string{"first", "second", "third", "final", "third", "second", "first"}
	if !reflect.DeepEqual(results, expected) {
		t.Fatalf("expected %v but got %v", expected, results)
	}
}
