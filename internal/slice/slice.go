//go:generate go get github.com/willfaught/goo/cmd/goo
//go:generate goo -in slice.go -out ../../slice_bool.go -json {"name":"Bool","sort":false,"element":"bool"}
//go:generate goo -in slice.go -out ../../slice_byte.go -json {"name":"Byte","sort":true,"element":"byte"}
//go:generate goo -in slice.go -out ../../slice_complex64.go -json {"name":"Complex64","sort":false,"element":"complex64"}
//go:generate goo -in slice.go -out ../../slice_complex128.go -json {"name":"Complex128","sort":false,"element":"complex128"}
//go:generate goo -in slice.go -out ../../slice_float32.go -json {"name":"Float32","sort":true,"element":"float32"}
//go:generate goo -in slice.go -out ../../slice_float64.go -json {"name":"Float64","sort":true,"element":"float64"}
//go:generate goo -in slice.go -out ../../slice_int.go -json {"name":"Int","sort":true,"element":"int"}
//go:generate goo -in slice.go -out ../../slice_interface.go -json {"name":"Interface","sort":false,"element":"interface{}"}
//go:generate goo -in slice.go -out ../../slice_int8.go -json {"name":"Int8","sort":true,"element":"int8"}
//go:generate goo -in slice.go -out ../../slice_int16.go -json {"name":"Int16","sort":true,"element":"int16"}
//go:generate goo -in slice.go -out ../../slice_int32.go -json {"name":"Int32","sort":true,"element":"int32"}
//go:generate goo -in slice.go -out ../../slice_int64.go -json {"name":"Int64","sort":true,"element":"int64"}
//go:generate goo -in slice.go -out ../../slice_rune.go -json {"name":"Rune","sort":true,"element":"rune"}
//go:generate goo -in slice.go -out ../../slice_string.go -json {"name":"String","sort":true,"element":"string"}
//go:generate goo -in slice.go -out ../../slice_uint.go -json {"name":"Uint","sort":true,"element":"uint"}
//go:generate goo -in slice.go -out ../../slice_uintptr.go -json {"name":"Uintptr","sort":true,"element":"uintptr"}
//go:generate goo -in slice.go -out ../../slice_uint8.go -json {"name":"Uint8","sort":true,"element":"uint8"}
//go:generate goo -in slice.go -out ../../slice_uint16.go -json {"name":"Uint16","sort":true,"element":"uint16"}
//go:generate goo -in slice.go -out ../../slice_uint32.go -json {"name":"Uint32","sort":true,"element":"uint32"}
//go:generate goo -in slice.go -out ../../slice_uint64.go -json {"name":"Uint64","sort":true,"element":"uint64"}

package goo

/// {{if .sort}}
import "sort" /// {{end}}

/// {{if false}}
type Slice interface{} /// {{end}}

/// {{if false}}
type __element__ int /// {{end}}

// Slice__name__ is a slice of __element__.
type Slice__name__ []__element__

var _ Slice = Slice__name__(nil)

/// {{if .sort}}
var _ sort.Interface = Slice__name__(nil) /// {{end}}

// Append appends v to s and returns the result.
func (s Slice__name__) Append(v ...interface{}) Slice {
	for _, v := range v {
		s = append(s, v.(__element__))
	}

	return s
}

// AppendSlice appends t to s and returns the result.
func (s Slice__name__) AppendSlice(t Slice) Slice {
	return append(s, t.(Slice__name__)...)
}

// Cap returns the s capacity.
func (s Slice__name__) Cap() int {
	return cap(s)
}

// Copy copies t to s.
func (s Slice__name__) Copy(t Slice) int {
	return copy(s, t.(Slice__name__))
}

// Equals returns whether s equals v.
func (s Slice__name__) Equals(v interface{}) bool {
	var t = v.(Slice__name__)
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
func (s Slice__name__) Get(i int) interface{} {
	return s[i]
}

// Len returns the s length.
func (s Slice__name__) Len() int {
	return len(s)
}

/// {{if .sort}}
// Less implements sort.Interface.
func (s Slice__name__) Less(i, j int) bool {
	return s[i] < s[j]
} /// {{end}}

// Make returns a new Slice__name__ with length l and capacity c.
func (s Slice__name__) Make(l, c int) Slice {
	return make(Slice__name__, l, c)
}

// Set sets the s element at index i to v.
func (s Slice__name__) Set(i int, v interface{}) {
	s[i] = v.(__element__)
}

// Slice returns the slice of s from indexes i to j.
func (s Slice__name__) Slice(i, j int) Slice {
	return s[i:j]
}

// SliceCap returns the slice of s from indexes i to j with capacity c.
func (s Slice__name__) SliceCap(i, j, c int) Slice {
	return s[i:j:c]
}

/// {{if .sort}}
// Swap implements sort.Interface.
func (s Slice__name__) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
} /// {{end}}
