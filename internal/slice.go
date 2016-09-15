//go:generate go install github.com/willfaught/goo/cmd/goo
//go:generate goo -in slice.go -out bool.go -var Name=Bool -var Type=bool -var Zero=false

package slice

import (
	"sort"

	"github.com/willfaught/goo/slice"
)

type __Type__ int //goo:omit

var __Zero__ = 0 //goo:omit

// Slice__Name__ is a slice of __Type__.
type Slice__Name__ []__Type__

// TODO: var _ goo.Equatable = Slice__Name__(nil)

/// {{if eq .Type "bool" "interface{}" | not}}
var _ sort.Interface = Slice__Name__(nil) /// {{end}}

// Append appends v to s and returns the result.
func (s Slice__Name__) Append(v ...interface{}) slice.Slice {
	for _, v := range v {
		s = append(s, v.(__Type__))
	}

	return s
}

// AppendSlice appends t to s and returns the result.
func (s Slice__Name__) AppendSlice(t slice.Slice) slice.Slice {
	return append(s, t.(Slice__Name__)...)
}

// Cap returns the s capacity.
func (s Slice__Name__) Cap() int {
	return cap(s)
}

// Copy copies t to s.
func (s Slice__Name__) Copy(t slice.Slice) int {
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

//goo:ifnot elementType bool
//goo:ifnot elementType interface{}
// Less implements sort.Interface.
func (s Slice__Name__) Less(i, j int) bool {
	return s[i] < s[j]
}

// Make returns a new slice.Slice__Name__ with length l and capacity c.
func (s Slice__Name__) Make(l, c int) slice.Slice {
	return make(Slice__Name__, l, c)
}

// Set sets the s element at index i to v.
func (s Slice__Name__) Set(i int, v interface{}) {
	s[i] = v.(__Type__)
}

// Slice returns the slice of s from indexes i to j.
func (s Slice__Name__) Slice(i, j int) slice.Slice {
	return s[i:j]
}

// SliceCap returns the slice of s from indexes i to j with capacity c.
func (s Slice__Name__) SliceCap(i, j, c int) slice.Slice {
	return s[i:j:c]
}

// {{if eq .Type "bool" "interface{}" | not}}
// Swap implements sort.Interface.
func (s Slice__Name__) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
} // {{end}}

// Zero returns the zero value of the s element type.
func (s Slice__Name__) Zero() interface{} {
	return __Zero__
}
