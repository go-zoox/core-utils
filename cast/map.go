package cast

import (
	"fmt"
	"reflect"
)

// ToMap casts a struct to map.
// reference: https://www.sobyte.net/post/2021-06/several-ways-to-convert-struct-to-mapstringinterface/
func ToMap(in interface{}, tagName ...string) (map[string]interface{}, error) {
	tagNameX := ""
	if len(tagName) > 0 && tagName[0] != "" {
		tagNameX = tagName[0]
	}

	out := make(map[string]interface{})

	v := reflect.ValueOf(in)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		return nil, fmt.Errorf("ToMap only accepts struct or struct point, but got %T", v)
	}

	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		fi := t.Field(i)
		vi := v.Field(i)

		key := fi.Name
		value := vi.Interface()
		isValueNestStruct := false

		if vi.Kind() == reflect.Ptr {
			vi = vi.Elem()
		}

		if vi.Kind() == reflect.Struct {
			isValueNestStruct = true
		}

		if tagNameX != "" {
			if tagValue := fi.Tag.Get(tagNameX); tagValue != "" {
				key = tagValue
			}
		}

		if !isValueNestStruct {
			out[key] = value
		} else {
			var err error
			out[key], err = ToMap(value, tagName...)
			if err != nil {
				return nil, err
			}
		}
	}

	return out, nil
}
