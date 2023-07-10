package reflect

import (
	goreflect "reflect"
)

// IsNil ...
func IsNil(v goreflect.Value) bool {
	switch v.Kind() {
	case goreflect.Chan, goreflect.Func, goreflect.Map, goreflect.Ptr, goreflect.Interface, goreflect.Slice:
		return v.IsNil()
	default:
		return false
	}
}
