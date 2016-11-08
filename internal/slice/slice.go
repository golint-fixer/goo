//go:generate go get github.com/willfaught/goo/cmd/goo
//go:generate goo macro slice.go ../../slice_bool.go -d {"Name":"Bool","Sort":false,"Element":"bool"}
//go:generate goo macro slice.go ../../slice_byte.go -d {"Name":"Byte","Sort":true,"Element":"byte"}
//go:generate goo macro slice.go ../../slice_complex64.go -d {"Name":"Complex64","Sort":false,"Element":"complex64"}
//go:generate goo macro slice.go ../../slice_complex128.go -d {"Name":"Complex128","Sort":false,"Element":"complex128"}
//go:generate goo macro slice.go ../../slice_float32.go -d {"Name":"Float32","Sort":true,"Element":"float32"}
//go:generate goo macro slice.go ../../slice_float64.go -d {"Name":"Float64","Sort":true,"Element":"float64"}
//go:generate goo macro slice.go ../../slice_int.go -d {"Name":"Int","Sort":true,"Element":"int"}
//go:generate goo macro slice.go ../../slice_interface.go -d {"Name":"Interface","Sort":false,"Element":"interface{}"}
//go:generate goo macro slice.go ../../slice_int8.go -d {"Name":"Int8","Sort":true,"Element":"int8"}
//go:generate goo macro slice.go ../../slice_int16.go -d {"Name":"Int16","Sort":true,"Element":"int16"}
//go:generate goo macro slice.go ../../slice_int32.go -d {"Name":"Int32","Sort":true,"Element":"int32"}
//go:generate goo macro slice.go ../../slice_int64.go -d {"Name":"Int64","Sort":true,"Element":"int64"}
//go:generate goo macro slice.go ../../slice_rune.go -d {"Name":"Rune","Sort":true,"Element":"rune"}
//go:generate goo macro slice.go ../../slice_string.go -d {"Name":"String","Sort":true,"Element":"string"}
//go:generate goo macro slice.go ../../slice_uint.go -d {"Name":"Uint","Sort":true,"Element":"uint"}
//go:generate goo macro slice.go ../../slice_uintptr.go -d {"Name":"Uintptr","Sort":true,"Element":"uintptr"}
//go:generate goo macro slice.go ../../slice_uint8.go -d {"Name":"Uint8","Sort":true,"Element":"uint8"}
//go:generate goo macro slice.go ../../slice_uint16.go -d {"Name":"Uint16","Sort":true,"Element":"uint16"}
//go:generate goo macro slice.go ../../slice_uint32.go -d {"Name":"Uint32","Sort":true,"Element":"uint32"}
//go:generate goo macro slice.go ../../slice_uint64.go -d {"Name":"Uint64","Sort":true,"Element":"uint64"}

package goo

/// {{if .Sort}}
import "sort" /// {{end}}

/// {{if false}}
type Slice interface{} /// {{end}}

/// {{if false}}
type __FIELD_Element__ int /// {{end}}

// Slice__FIELD_Name__ is a slice of __FIELD_Element__.
type Slice__FIELD_Name__ []__FIELD_Element__

var _ Slice = Slice__FIELD_Name__(nil)

/// {{if .Sort}}
var _ sort.Interface = Slice__FIELD_Name__(nil) /// {{end}}

// Append appends v to s and returns the result.
func (s Slice__FIELD_Name__) Append(v ...interface{}) Slice {
	for _, v := range v {
		s = append(s, v.(__FIELD_Element__))
	}

	return s
}

// AppendSlice appends t to s and returns the result.
func (s Slice__FIELD_Name__) AppendSlice(t Slice) Slice {
	return append(s, t.(Slice__FIELD_Name__)...)
}

// Cap returns the s capacity.
func (s Slice__FIELD_Name__) Cap() int {
	return cap(s)
}

// Copy copies t to s.
func (s Slice__FIELD_Name__) Copy(t Slice) int {
	return copy(s, t.(Slice__FIELD_Name__))
}

// Equals returns whether s equals v.
func (s Slice__FIELD_Name__) Equals(v interface{}) bool {
	var t = v.(Slice__FIELD_Name__)
	var l = len(s)

	if len(t) != l {
		return false
	}

	if l == 0 {
		return true
	}

	for i := range s {
		if t[i] != s[i] {
			return false
		}
	}

	return true
}

// Get returns the s element at index i.
func (s Slice__FIELD_Name__) Get(i int) interface{} {
	return s[i]
}

// GetRange returns the slice of s from indexes i to j.
func (s Slice__FIELD_Name__) GetRange(i, j int) Slice {
	return s[i:j]
}

// GetRangeCap returns the slice of s from indexes i to j with capacity c.
func (s Slice__FIELD_Name__) GetRangeCap(i, j, c int) Slice {
	return s[i:j:c]
}

// Len returns the s length.
func (s Slice__FIELD_Name__) Len() int {
	return len(s)
}

/// {{if .Sort}}
// Less implements sort.Interface.
func (s Slice__FIELD_Name__) Less(i, j int) bool {
	return s[i] < s[j]
} /// {{end}}

// Make returns a new Slice__FIELD_Name__ with length l and capacity c.
func (s Slice__FIELD_Name__) Make(l, c int) Slice {
	return make(Slice__FIELD_Name__, l, c)
}

// Set sets the s element at index i to v.
func (s Slice__FIELD_Name__) Set(i int, v interface{}) {
	s[i] = v.(__FIELD_Element__)
}

/// {{if .Sort}}
// Swap implements sort.Interface.
func (s Slice__FIELD_Name__) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
} /// {{end}}
