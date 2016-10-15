//go:generate go get github.com/willfaught/goo/cmd/goo
//go:generate goo macro -i slice.go -o ../../slice_bool.go -j {"name":"Bool","sort":false,"element":"bool"}
//go:generate goo macro -i slice.go -o ../../slice_byte.go -j {"name":"Byte","sort":true,"element":"byte"}
//go:generate goo macro -i slice.go -o ../../slice_complex64.go -j {"name":"Complex64","sort":false,"element":"complex64"}
//go:generate goo macro -i slice.go -o ../../slice_complex128.go -j {"name":"Complex128","sort":false,"element":"complex128"}
//go:generate goo macro -i slice.go -o ../../slice_float32.go -j {"name":"Float32","sort":true,"element":"float32"}
//go:generate goo macro -i slice.go -o ../../slice_float64.go -j {"name":"Float64","sort":true,"element":"float64"}
//go:generate goo macro -i slice.go -o ../../slice_int.go -j {"name":"Int","sort":true,"element":"int"}
//go:generate goo macro -i slice.go -o ../../slice_interface.go -j {"name":"Interface","sort":false,"element":"interface{}"}
//go:generate goo macro -i slice.go -o ../../slice_int8.go -j {"name":"Int8","sort":true,"element":"int8"}
//go:generate goo macro -i slice.go -o ../../slice_int16.go -j {"name":"Int16","sort":true,"element":"int16"}
//go:generate goo macro -i slice.go -o ../../slice_int32.go -j {"name":"Int32","sort":true,"element":"int32"}
//go:generate goo macro -i slice.go -o ../../slice_int64.go -j {"name":"Int64","sort":true,"element":"int64"}
//go:generate goo macro -i slice.go -o ../../slice_rune.go -j {"name":"Rune","sort":true,"element":"rune"}
//go:generate goo macro -i slice.go -o ../../slice_string.go -j {"name":"String","sort":true,"element":"string"}
//go:generate goo macro -i slice.go -o ../../slice_uint.go -j {"name":"Uint","sort":true,"element":"uint"}
//go:generate goo macro -i slice.go -o ../../slice_uintptr.go -j {"name":"Uintptr","sort":true,"element":"uintptr"}
//go:generate goo macro -i slice.go -o ../../slice_uint8.go -j {"name":"Uint8","sort":true,"element":"uint8"}
//go:generate goo macro -i slice.go -o ../../slice_uint16.go -j {"name":"Uint16","sort":true,"element":"uint16"}
//go:generate goo macro -i slice.go -o ../../slice_uint32.go -j {"name":"Uint32","sort":true,"element":"uint32"}
//go:generate goo macro -i slice.go -o ../../slice_uint64.go -j {"name":"Uint64","sort":true,"element":"uint64"}

package goo

/// {{if .sort}}
import "sort" /// {{end}}

/// {{if false}}
type Slice interface{} /// {{end}}

/// {{if false}}
type __FIELD_element__ int /// {{end}}

// Slice__FIELD_name__ is a slice of __FIELD_element__.
type Slice__FIELD_name__ []__FIELD_element__

var _ Slice = Slice__FIELD_name__(nil)

/// {{if .sort}}
var _ sort.Interface = Slice__FIELD_name__(nil) /// {{end}}

// Append appends v to s and returns the result.
func (s Slice__FIELD_name__) Append(v ...interface{}) Slice {
	for _, v := range v {
		s = append(s, v.(__FIELD_element__))
	}

	return s
}

// AppendSlice appends t to s and returns the result.
func (s Slice__FIELD_name__) AppendSlice(t Slice) Slice {
	return append(s, t.(Slice__FIELD_name__)...)
}

// Cap returns the s capacity.
func (s Slice__FIELD_name__) Cap() int {
	return cap(s)
}

// Copy copies t to s.
func (s Slice__FIELD_name__) Copy(t Slice) int {
	return copy(s, t.(Slice__FIELD_name__))
}

// Equals returns whether s equals v.
func (s Slice__FIELD_name__) Equals(v interface{}) bool {
	var t = v.(Slice__FIELD_name__)
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
func (s Slice__FIELD_name__) Get(i int) interface{} {
	return s[i]
}

// Len returns the s length.
func (s Slice__FIELD_name__) Len() int {
	return len(s)
}

/// {{if .sort}}
// Less implements sort.Interface.
func (s Slice__FIELD_name__) Less(i, j int) bool {
	return s[i] < s[j]
} /// {{end}}

// Make returns a new Slice__FIELD_name__ with length l and capacity c.
func (s Slice__FIELD_name__) Make(l, c int) Slice {
	return make(Slice__FIELD_name__, l, c)
}

// Set sets the s element at index i to v.
func (s Slice__FIELD_name__) Set(i int, v interface{}) {
	s[i] = v.(__FIELD_element__)
}

// Slice returns the slice of s from indexes i to j.
func (s Slice__FIELD_name__) Slice(i, j int) Slice {
	return s[i:j]
}

// SliceCap returns the slice of s from indexes i to j with capacity c.
func (s Slice__FIELD_name__) SliceCap(i, j, c int) Slice {
	return s[i:j:c]
}

/// {{if .sort}}
// Swap implements sort.Interface.
func (s Slice__FIELD_name__) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
} /// {{end}}
