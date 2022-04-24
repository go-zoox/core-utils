package object

import (
	"fmt"
	"reflect"
)

// ToMap transform struct to an may.
func ToMap(in any, tag ...string) (map[string]any, error) {
	var tagX string
	if len(tag) > 0 {
		tagX = tag[0]
	}

	out := make(map[string]any)

	v := reflect.ValueOf(in)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		return nil, fmt.Errorf("ToMap only accepts struct, but got %T", in)
	}

	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		key := t.Field(i)
		value := v.Field(i)

		switch value.Type().Kind() {
		case reflect.Ptr:
			value = value.Elem()
			fallthrough
		case reflect.Struct:
			inValue, err := ToMap(value.Interface(), tagX)
			if err != nil {
				return nil, err
			}

			if tagX != "" {
				if tagValue := key.Tag.Get(tagX); tagValue != "" {
					out[tagValue] = inValue
				}
			} else {
				out[key.Name] = inValue
			}
		default:
			if tagX != "" {
				if tagValue := key.Tag.Get(tagX); tagValue != "" {
					out[tagValue] = value.Interface()
				}
			} else {
				out[key.Name] = value.Interface()
			}
		}
	}

	return out, nil
}
