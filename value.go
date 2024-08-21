package huma

import "reflect"

type Value interface {
	Kind() reflect.Kind
	Type() reflect.Type
	IsNil() bool
	IsValid() bool

	Addr() reflect.Value
	CanAddr() bool
	Len() int
	Index(i int) reflect.Value

	IsZero() bool
	Bool() bool
	Bytes() []byte
	Float() float64
	Int() int64
	Uint() uint64
	Interface() (i any)
	String() string

	SetZero()
	SetBool(x bool)
	SetBytes(x []byte)
	SetFloat(x float64)
	SetInt(x int64)
	SetUint(x uint64)
	SetInterface(i any)
	SetString(x string)
	Set(x reflect.Value)

	Value() reflect.Value
}

type value struct {
	v reflect.Value
}

func (v *value) Value() reflect.Value {
	return v.instantiate()
}

func (v *value) Kind() reflect.Kind {
	return reflectValueKind(v.v)
}

func (v *value) Type() reflect.Type {
	return reflectValueType(v.v)
}

func (v *value) IsNil() bool {
	return reflectValueIsNil(v.v)
}

func (v *value) IsValid() bool {
	return reflectValueElem(v.v).IsValid()
}

func (v *value) IsZero() bool {
	return reflectValueElem(v.v).IsZero()
}

func (v *value) Bool() bool {
	return reflectValueElem(v.v).Bool()
}

func (v *value) Bytes() []byte {
	return reflectValueElem(v.v).Bytes()
}

func (v *value) Float() float64 {
	return reflectValueElem(v.v).Float()
}

func (v *value) Int() int64 {
	return reflectValueElem(v.v).Int()
}

func (v *value) Uint() uint64 {
	return reflectValueElem(v.v).Uint()
}

func (v *value) Interface() (i any) {
	return reflectValueElem(v.v).Interface()
}

func (v *value) String() string {
	return reflectValueElem(v.v).String()
}

func (v *value) CanAddr() bool {
	return reflectValueElem(v.v).CanAddr()
}

func (v *value) Addr() reflect.Value {
	return reflectValueElem(v.v).Addr()
}

func (v *value) Len() int {
	return v.v.Len()
}

func (v *value) Index(i int) reflect.Value {
	return v.v.Index(i)
}

func (v *value) SetZero() {
	v.v.SetZero()
}

func (v *value) instantiate() reflect.Value {
	vi := v.v
	for vi.Kind() == reflect.Ptr {
		if vi.IsNil() {
			vi.Set(reflect.New(vi.Type().Elem()))
		}
		vi = vi.Elem()
	}
	return vi
}

func (v *value) SetBool(x bool) {
	v.instantiate().SetBool(x)
}

func (v *value) SetBytes(x []byte) {
	v.instantiate().SetBytes(x)
}

func (v *value) SetFloat(x float64) {
	v.instantiate().SetFloat(x)
}

func (v *value) SetInt(x int64) {
	v.instantiate().SetInt(x)
}

func (v *value) SetUint(x uint64) {
	v.instantiate().SetUint(x)
}

func (v *value) SetInterface(i any) {
	v.instantiate().Set(reflect.ValueOf(i))
}

func (v *value) SetString(x string) {
	v.instantiate().SetString(x)
}

func (v *value) Set(x reflect.Value) {
	v.instantiate().Set(x)
}

// reflectValueIsNil returns the first non-pointer type from the [reflect.Value].
func reflectValueIsNil(v reflect.Value) bool {
	for v.Kind() == reflect.Ptr {
		if v.IsNil() {
			return true
		}
		v = v.Elem()
	}
	return false
}

// reflectValueElem returns the first non-pointer type from the [reflect.Value].
func reflectValueElem(v reflect.Value) reflect.Value {
	for v.Kind() == reflect.Ptr {
		if v.IsNil() {
			return reflectValueZero(v)
		}
		v = v.Elem()
	}
	return v
}

func reflectValueKind(v reflect.Value) reflect.Kind {
	t := v.Type()
	for t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	return t.Kind()
}

func reflectValueType(v reflect.Value) reflect.Type {
	t := v.Type()
	for t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	return t
}

func reflectValueZero(v reflect.Value) reflect.Value {
	t := v.Type()
	for t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	return reflect.Zero(t)
}
