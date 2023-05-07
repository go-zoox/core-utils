package object

import (
	"reflect"
	"regexp"
	gostrings "strings"

	"github.com/go-zoox/core-utils/cast"
)

// Get returns the value of the key in the object.
// Support key path
//
// Example:
//
//		object.Get(map[string]interface{}{"a": 1}, "a")
//	 // 1
//
//		object.Get(map[string]interface{}{"a": map[string]interface{}{"b": 2}}, "a.b")
//		// 2
func Get[K comparable, V any](object map[K]V, key string) V {
	keysX := gostrings.Split(key, ".")
	var keys []K
	for _, keyX := range keysX {
		keys = append(keys, reflect.ValueOf(keyX).Interface().(K))
	}

	keyLength := len(keys)
	tmp := reflect.ValueOf(object).Interface()
	for index, k := range keys {
		// array
		kt := reflect.ValueOf(k).Interface().(string)
		// ttt, _ := regexp.MatchString("^[0-9]+$", kt)
		// fmt.Println("xxx:", k, kt, ttt, tmp)
		if ok := isNumeric(kt); ok {
			if reflect.TypeOf(tmp).Kind() != reflect.Slice {
				var empty V
				return empty
			}

			ktt := cast.ToInt(kt)
			if ktt >= reflect.ValueOf(tmp).Len() {
				var empty V
				return empty
			}

			v := reflect.ValueOf(tmp).Index(ktt).Interface()
			if index == keyLength-1 {
				return reflect.ValueOf(v).Interface().(V)
			}

			tmp = v
			continue
		}

		if v, ok := tmp.(map[K]V)[k]; ok {
			if index == keyLength-1 {
				return reflect.ValueOf(v).Interface().(V)
			}

			switch reflect.TypeOf(v).Kind() {
			case reflect.Map:
				tmp = reflect.ValueOf(v).Interface().(map[K]V)
			case reflect.Slice:
				tmp = reflect.ValueOf(v).Interface()
			default:
				tmp = v
			}

			// fmt.Println("key:", k, tmp)
		} else {
			var empty V
			return empty
		}
	}

	var empty V
	return empty
}

func isNumeric(s string) bool {
	return numericRegExp.MatchString(s)
}

var numericRegExp, _ = regexp.Compile("^[0-9]+$")
