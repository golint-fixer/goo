//go:generate go get github.com/willfaught/goo/cmd/goo
//go:generate goo macro slices.go ../../slice_bool.go -d {"Element":"bool","Name":"Bool","Sort":false}
//go:generate goo macro slices.go ../../slice_byte.go -d {"Element":"byte","Name":"Byte","Sort":true}
//go:generate goo macro slices.go ../../slice_complex64.go -d {"Element":"complex64","Name":"Complex64","Sort":false}
//go:generate goo macro slices.go ../../slice_complex128.go -d {"Element":"complex128","Name":"Complex128","Sort":false}
//go:generate goo macro slices.go ../../slice_float32.go -d {"Element":"float32","Name":"Float32","Sort":true}
//go:generate goo macro slices.go ../../slice_float64.go -d {"Element":"float64","Name":"Float64","Sort":true}
//go:generate goo macro slices.go ../../slice_int.go -d {"Element":"int","Name":"Int","Sort":true}
//go:generate goo macro slices.go ../../slice_interface.go -d {"Element":"interface{}","Name":"Interface","Sort":false}
//go:generate goo macro slices.go ../../slice_int8.go -d {"Element":"int8","Name":"Int8","Sort":true}
//go:generate goo macro slices.go ../../slice_int16.go -d {"Element":"int16","Name":"Int16","Sort":true}
//go:generate goo macro slices.go ../../slice_int32.go -d {"Element":"int32","Name":"Int32","Sort":true}
//go:generate goo macro slices.go ../../slice_int64.go -d {"Element":"int64","Name":"Int64","Sort":true}
//go:generate goo macro slices.go ../../slice_rune.go -d {"Element":"rune","Name":"Rune","Sort":true}
//go:generate goo macro slices.go ../../slice_string.go -d {"Element":"string","Name":"String","Sort":true}
//go:generate goo macro slices.go ../../slice_uint.go -d {"Element":"uint","Name":"Uint","Sort":true}
//go:generate goo macro slices.go ../../slice_uintptr.go -d {"Element":"uintptr","Name":"Uintptr","Sort":true}
//go:generate goo macro slices.go ../../slice_uint8.go -d {"Element":"uint8","Name":"Uint8","Sort":true}
//go:generate goo macro slices.go ../../slice_uint16.go -d {"Element":"uint16","Name":"Uint16","Sort":true}
//go:generate goo macro slices.go ../../slice_uint32.go -d {"Element":"uint32","Name":"Uint32","Sort":true}
//go:generate goo macro slices.go ../../slice_uint64.go -d {"Element":"uint64","Name":"Uint64","Sort":true}

package goo

/// {{if .Sort}}
import "sort" /// {{end}}

/// {{if false}}
type __FIELD_Element__ int
type Equatable interface{}
type Pointer interface{}
type Slice interface{}
type Value interface{} /// {{end}}

// Slice__FIELD_Name__Zero is the Slice__FIELD_Name__ zero value.
var Slice__FIELD_Name__Zero = Slice__FIELD_Name__(nil)

var _ Pointer = &Slice__FIELD_Name__{}

var _ Slice = Slice__FIELD_Name__(nil)

/// {{if .Sort}}
var _ sort.Interface = Slice__FIELD_Name__(nil) /// {{end}}

// Slice__FIELD_Name__ is a slice of __FIELD_Element__.
type Slice__FIELD_Name__ []__FIELD_Element__

// Append implements Slice.
func (s Slice__FIELD_Name__) Append(v ...interface{}) Slice {
	for _, v := range v {
		s = append(s, v.(__FIELD_Element__))
	}

	return s
}

// AppendSlice implements Slice.
func (s Slice__FIELD_Name__) AppendSlice(other Slice) Slice {
	return append(s, other.(Slice__FIELD_Name__)...)
}

// Cap implements Slice.
func (s Slice__FIELD_Name__) Cap() int {
	return cap(s)
}

// Copy implements Slice.
func (s Slice__FIELD_Name__) Copy(other Slice) int {
	return copy(s, other.(Slice__FIELD_Name__))
}

// Dereference implements Slice.
func (s *Slice__FIELD_Name__) Dereference() Value {
	return *s
}

// Equals implements Slice.
func (s Slice__FIELD_Name__) Equals(other Equatable) bool {
	var t = other.(Slice__FIELD_Name__)

	if len(t) != len(s) {
		return false
	}

	for i := range s {
		if t[i] != s[i] {
			return false
		}
	}

	return true
}

// Get implements Slice.
func (s Slice__FIELD_Name__) Get(i int) interface{} {
	return s[i]
}

// GetRange implements Slice.
func (s Slice__FIELD_Name__) GetRange(i, j int) Slice {
	return s[i:j]
}

// GetRangeCap implements Slice.
func (s Slice__FIELD_Name__) GetRangeCap(i, j, c int) Slice {
	return s[i:j:c]
}

// Len implements Slice.
func (s Slice__FIELD_Name__) Len() int {
	return len(s)
}

/// {{if .Sort}}
// Less implements sort.Interface.
func (s Slice__FIELD_Name__) Less(i, j int) bool {
	return s[i] < s[j]
} /// {{end}}

// Make implements Slice.
func (s Slice__FIELD_Name__) Make(l, c int) Slice {
	return make(Slice__FIELD_Name__, l, c)
}

// NotEquals implements Slice.
func (s Slice__FIELD_Name__) NotEquals(other Equatable) bool {
	return !s.Equals(other)
}

// Reference implements Slice.
func (s Slice__FIELD_Name__) Reference() Pointer {
	return &s
}

// Set implements Slice.
func (s Slice__FIELD_Name__) Set(i int, v interface{}) {
	s[i] = v.(__FIELD_Element__)
}

/// {{if .Sort}}
// Swap implements sort.Interface.
func (s Slice__FIELD_Name__) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
} /// {{end}}
