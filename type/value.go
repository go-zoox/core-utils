package typ

import "fmt"

// Value ...
type Value struct {
	Data map[string]interface{}
	key  string
}

// NewValue creates a new Value.
func NewValue(v ...map[string]interface{}) *Value {
	vx := map[string]interface{}{}
	if len(v) > 0 && v[0] != nil {
		vx = v[0]
	}

	return &Value{
		Data: vx,
	}
}

// Get gets the value by key.
func (v *Value) Get(key string) *Value {
	return &Value{
		Data: v.Data,
		key:  key,
	}
}

// Set sets the value by key.
func (v *Value) Set(key string, value interface{}) {
	v.Data[key] = value
}

// Has checks if the key has been set.
func (v *Value) Has(key string) bool {
	_, ok := v.Data[key]
	return ok
}

// String ...
func (v *Value) String() string {
	if v.key == "" {
		panic("call Get(key) first")
	}

	if v, ok := v.Data[v.key].(string); ok {
		return v
	}

	return ""
}

// Int64 ...
func (v *Value) Int64() int64 {
	if v.key == "" {
		panic("call Get(key) first")
	}

	if v, ok := v.Data[v.key].(int64); ok {
		return v
	}

	return int64(v.Float64())
}

// Int64E ...
func (v *Value) Int64E() (vv int64, err error) {
	if v.key == "" {
		panic("call Get(key) first")
	}

	if v, ok := v.Data[v.key].(int64); ok {
		return v, nil
	}

	return 0, fmt.Errorf("value(%v) cannot be translated to int64", v.Data[v.key])
}

// Int ...
func (v *Value) Int() int {
	if v.key == "" {
		panic("call Get(key) first")
	}

	if v, ok := v.Data[v.key].(int); ok {
		return v
	}

	return 0
}

// IntE ...
func (v *Value) IntE() (vv int, err error) {
	if v.key == "" {
		panic("call Get(key) first")
	}

	if v, ok := v.Data[v.key].(int); ok {
		return v, nil
	}

	return 0, fmt.Errorf("value(%v) cannot be translated to int", v.Data[v.key])
}

// Bool ...
func (v *Value) Bool() bool {
	if v.key == "" {
		panic("call Get(key) first")
	}

	if v, ok := v.Data[v.key].(bool); ok {
		return v
	}

	return false
}

// BoolE ...
func (v *Value) BoolE() (vv bool, err error) {
	if v.key == "" {
		panic("call Get(key) first")
	}

	if v, ok := v.Data[v.key].(bool); ok {
		return v, nil
	}

	return false, fmt.Errorf("value(%v) cannot be translated to bool", v.Data[v.key])
}

// Float64 ...
func (v *Value) Float64() float64 {
	if v.key == "" {
		panic("call Get(key) first")
	}

	if v, ok := v.Data[v.key].(float64); ok {
		return v
	}

	return 0
}

// Float64E ...
func (v *Value) Float64E() (vv float64, err error) {
	if v.key == "" {
		panic("call Get(key) first")
	}

	if v, ok := v.Data[v.key].(float64); ok {
		return v, nil
	}

	return 0, fmt.Errorf("value(%v) cannot be translated to float64", v.Data[v.key])
}
