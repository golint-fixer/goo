//go:generate go get github.com/willfaught/goo/cmd/goo
//go:generate goo -in slice.go -out ../../slice_bool.go -json {"Name":"Bool","Sort":fasle,"Type":"bool","Zero":"false"}
//go:generate goo -in slice.go -out ../../slice_byte.go -json {"Name":"Byte","Sort":true,"Type":"byte","Zero":"0"}
//go:generate goo -in slice.go -out ../../slice_complex64.go -json {"Name":"Complex64","Sort":false,"Type":"complex64","Zero":"0"}
//go:generate goo -in slice.go -out ../../slice_complex128.go -json {"Name":"Complex128","Sort":false,"Type":"complex128","Zero":"0"}
//go:generate goo -in slice.go -out ../../slice_float32.go -json {"Name":"Float32","Sort":true,"Type":"float32","Zero":"0"}
//go:generate goo -in slice.go -out ../../slice_float64.go -json {"Name":"Float64","Sort":true,"Type":"float64","Zero":"0"}
//go:generate goo -in slice.go -out ../../slice_int.go -json {"Name":"Int","Sort":true,"Type":"int","Zero":"0"}
//go:generate goo -in slice.go -out ../../slice_interface.go -json {"Name":"Interface","Sort":false,"Type":"interface{}","Zero":"nil"}
//go:generate goo -in slice.go -out ../../slice_int8.go -json {"Name":"Int8","Sort":true,"Type":"int8","Zero":"0"}
//go:generate goo -in slice.go -out ../../slice_int16.go -json {"Name":"Int16","Sort":true,"Type":"int16","Zero":"0"}
//go:generate goo -in slice.go -out ../../slice_int32.go -json {"Name":"Int32","Sort":true,"Type":"int32","Zero":"0"}
//go:generate goo -in slice.go -out ../../slice_int64.go -json {"Name":"Int64","Sort":true,"Type":"int64","Zero":"0"}
//go:generate goo -in slice.go -out ../../slice_rune.go -json {"Name":"Rune","Sort":true,"Type":"rune","Zero":"0"}
//go:generate goo -in slice.go -out ../../slice_string.go -json {"Name":"String","Sort":true,"Type":"string","Zero":"\"\""}
//go:generate goo -in slice.go -out ../../slice_uint.go -json {"Name":"Uint","Sort":true,"Type":"uint","Zero":"0"}
//go:generate goo -in slice.go -out ../../slice_uintptr.go -json {"Name":"Uintptr","Sort":true,"Type":"uintptr","Zero":"0"}
//go:generate goo -in slice.go -out ../../slice_uint8.go -json {"Name":"Uint8","Sort":true,"Type":"uint8","Zero":"0"}
//go:generate goo -in slice.go -out ../../slice_uint16.go -json {"Name":"Uint16","Sort":true,"Type":"uint16","Zero":"0"}
//go:generate goo -in slice.go -out ../../slice_uint32.go -json {"Name":"Uint32","Sort":true,"Type":"uint32","Zero":"0"}
//go:generate goo -in slice.go -out ../../slice_uint64.go -json {"Name":"Uint64","Sort":true,"Type":"uint64","Zero":"0"}

package goo

/// {{if .Sort}}
import "sort" /// {{end}}

/// {{if false}}
type Slice interface{} /// {{end}}

/// {{if false}}
type __Type__ int /// {{end}}

/// {{if false}}
var __Zero__ = 0 /// {{end}}

// Slice__Name__ is a slice of __Type__.
type Slice__Name__ []__Type__

/// {{if .Sort}}
var _ sort.Interface = Slice__Name__(nil) /// {{end}}

// Append appends v to s and returns the result.
func (s Slice__Name__) Append(v ...interface{}) Slice {
	for _, v := range v {
		s = append(s, v.(__Type__))
	}

	return s
}

// AppendSlice appends t to s and returns the result.
func (s Slice__Name__) AppendSlice(t Slice) Slice {
	return append(s, t.(Slice__Name__)...)
}

// Cap returns the s capacity.
func (s Slice__Name__) Cap() int {
	return cap(s)
}

// Copy copies t to s.
func (s Slice__Name__) Copy(t Slice) int {
	return copy(s, t.(Slice__Name__))
}

// Equals returns whether s equals v.
func (s Slice__Name__) Equals(v interface{}) bool {
	var t = v.(Slice__Name__)
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
func (s Slice__Name__) Get(i int) interface{} {
	return s[i]
}

// Len returns the s length.
func (s Slice__Name__) Len() int {
	return len(s)
}

/// {{if .Sort}}
// Less implements sort.Interface.
func (s Slice__Name__) Less(i, j int) bool {
	return s[i] < s[j]
} /// {{end}}

// Make returns a new Slice__Name__ with length l and capacity c.
func (s Slice__Name__) Make(l, c int) Slice {
	return make(Slice__Name__, l, c)
}

// Set sets the s element at index i to v.
func (s Slice__Name__) Set(i int, v interface{}) {
	s[i] = v.(__Type__)
}

// Slice returns the slice of s from indexes i to j.
func (s Slice__Name__) Slice(i, j int) Slice {
	return s[i:j]
}

// SliceCap returns the slice of s from indexes i to j with capacity c.
func (s Slice__Name__) SliceCap(i, j, c int) Slice {
	return s[i:j:c]
}

/// {{if .Sort}}
// Swap implements sort.Interface.
func (s Slice__Name__) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
} /// {{end}}

// Zero returns the zero __Zero__ value of the s element type.
func (s Slice__Name__) Zero() interface{} {
	return __Zero__
}
