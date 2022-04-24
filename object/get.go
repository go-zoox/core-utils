package object

import (
	"reflect"
	"strings"
)

// Get returns the value of the key in the object.
// Support key path
//
// Example:
// 	object.Get(map[string]interface{}{"a": 1}, "a")
//  // 1
//
//	object.Get(map[string]interface{}{"a": map[string]interface{}{"b": 2}}, "a.b")
//	// 2
func Get[K comparable, V any](object map[K]V, key string) V {
	keysX := strings.Split(key, ".")
	var keys []K
	for _, keyX := range keysX {
		keys = append(keys, reflect.ValueOf(keyX).Interface().(K))
	}

	keyLength := len(keys)
	tmp := object
	for index, k := range keys {
		// kt := reflect.ValueOf(k).Interface().(string)
		// if ok, err := regexp.MatchString("^[0-9]+$", kt); ok && err == nil {
		// 	ktt, _ := strconv.Atoi(kt)
		// 	k = reflect.ValueOf(ktt).Interface().(K)
		// }

		// fmt.Println("xxx:", k)
		if v, ok := tmp[k]; ok {
			if index == keyLength-1 {
				return reflect.ValueOf(v).Interface().(V)
			}

			switch reflect.ValueOf(v).Interface().(type) {
			case map[K]V:
				tmp = reflect.ValueOf(v).Interface().(map[K]V)
			case []V:
				// tmp = reflect.ValueOf(v).Interface()
			}
		} else {
			var empty V
			return empty
		}
	}

	var empty V
	return empty
}
