package object

import (
	"reflect"
	"strconv"
)

// Walk traverses object by deep first.
func Walk[K comparable, V any](object map[K]V, fn func(V, K, string)) {
	WalkByDeepFirst(object, fn, "")
}

// WalkByDeepFirst traverses object by deep first.
func WalkByDeepFirst[K comparable, V any](object map[K]V, fn func(V, K, string), parent string) {
	for key, value := range object {
		switch reflect.TypeOf(value).Kind() {
		case reflect.Map:
			if parent == "" {
				WalkByDeepFirst(reflect.ValueOf(value).Interface().(map[K]V), fn, reflect.ValueOf(key).Interface().(string))
			} else {
				WalkByDeepFirst(reflect.ValueOf(value).Interface().(map[K]V), fn, parent+"."+reflect.ValueOf(key).Interface().(string))
			}
		case reflect.Slice:
			for index, item := range reflect.ValueOf(value).Interface().([]V) {
				switch reflect.TypeOf(item).Kind() {
				case reflect.Map:
					if parent == "" {
						WalkByDeepFirst(reflect.ValueOf(item).Interface().(map[K]V), fn, reflect.ValueOf(key).Interface().(string)+"."+strconv.Itoa(index))
					} else {
						WalkByDeepFirst(reflect.ValueOf(item).Interface().(map[K]V), fn, parent+"."+reflect.ValueOf(key).Interface().(string)+"."+strconv.Itoa(index))
					}
				default:
					fn(value, key, "")
				}
			}
		default:
			if parent == "" {
				fn(value, key, reflect.ValueOf(key).Interface().(string))
			} else {
				fn(value, key, parent+"."+reflect.ValueOf(key).Interface().(string))
			}
		}
	}
}
