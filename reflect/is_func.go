package reflect

import (
	goreflect "reflect"
)

func IsFunc(val any) bool {
	if val == nil {
		return false
	}

	return goreflect.TypeOf(val).Kind() == goreflect.Func
}
