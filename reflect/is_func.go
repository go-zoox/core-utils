package reflect

import (
	goreflect "reflect"
)

// IsFunc ...
func IsFunc(val any) bool {
	if val == nil {
		return false
	}

	return goreflect.TypeOf(val).Kind() == goreflect.Func
}
