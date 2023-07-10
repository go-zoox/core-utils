package reflect

import (
	"reflect"
	goreflect "reflect"
	"testing"

	"github.com/go-zoox/testify"
)

func TestIsNil(t *testing.T) {
	testify.AssertFalse(t, IsNil(reflect.ValueOf(nil)))

	var v *goreflect.Value
	testify.AssertTrue(t, IsNil(reflect.ValueOf(v)))
}
