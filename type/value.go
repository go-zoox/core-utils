package typ

type Value struct {
	Data map[string]interface{}
	key  string
}

func NewValue(v ...map[string]interface{}) *Value {
	vx := map[string]interface{}{}
	if len(v) > 0 && v[0] != nil {
		vx = v[0]
	}

	return &Value{
		Data: vx,
	}
}

func (v *Value) Get(key string) *Value {
	return &Value{
		Data: v.Data,
		key:  key,
	}
}

func (v *Value) Set(key string, value interface{}) {
	v.Data[key] = value
}

func (v *Value) Has(key string) bool {
	_, ok := v.Data[key]
	return ok
}

func (v *Value) String() string {
	if v.key == "" {
		panic("call Get(key) first")
	}

	if v, ok := v.Data[v.key].(string); ok {
		return v
	}

	return ""
}

func (v *Value) Int64() int64 {
	if v.key == "" {
		panic("call Get(key) first")
	}

	if v, ok := v.Data[v.key].(int64); ok {
		return v
	}

	return int64(v.Float64())
}

func (v *Value) Int() int {
	if v.key == "" {
		panic("call Get(key) first")
	}

	if v, ok := v.Data[v.key].(int); ok {
		return v
	}

	return 0
}

func (v *Value) Bool() bool {
	if v.key == "" {
		panic("call Get(key) first")
	}

	if v, ok := v.Data[v.key].(bool); ok {
		return v
	}

	return false
}

func (v *Value) Float64() float64 {
	if v.key == "" {
		panic("call Get(key) first")
	}

	if v, ok := v.Data[v.key].(float64); ok {
		return v
	}

	return 0
}
