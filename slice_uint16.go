package goo

import "sort"

// SliceUint16 is a slice of uint16.
type SliceUint16 []uint16

var _ sort.Interface = SliceUint16(nil)

// Append appends v to s and returns the result.
func (s SliceUint16) Append(v ...interface{}) Slice {
	for _, v := range v {
		s = append(s, v.(uint16))
	}

	return s
}

// AppendSlice appends t to s and returns the result.
func (s SliceUint16) AppendSlice(t Slice) Slice {
	return append(s, t.(SliceUint16)...)
}

// Cap returns the s capacity.
func (s SliceUint16) Cap() int {
	return cap(s)
}

// Copy copies t to s.
func (s SliceUint16) Copy(t Slice) int {
	return copy(s, t.(SliceUint16))
}

// Equals returns whether s equals v.
func (s SliceUint16) Equals(v interface{}) bool {
	var t = v.(SliceUint16)
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
func (s SliceUint16) Get(i int) interface{} {
	return s[i]
}

// Len returns the s length.
func (s SliceUint16) Len() int {
	return len(s)
}

// Less implements sort.Interface.
func (s SliceUint16) Less(i, j int) bool {
	return s[i] < s[j]
}

// Make returns a new SliceUint16 with length l and capacity c.
func (s SliceUint16) Make(l, c int) Slice {
	return make(SliceUint16, l, c)
}

// Set sets the s element at index i to v.
func (s SliceUint16) Set(i int, v interface{}) {
	s[i] = v.(uint16)
}

// Slice returns the slice of s from indexes i to j.
func (s SliceUint16) Slice(i, j int) Slice {
	return s[i:j]
}

// SliceCap returns the slice of s from indexes i to j with capacity c.
func (s SliceUint16) SliceCap(i, j, c int) Slice {
	return s[i:j:c]
}

// Swap implements sort.Interface.
func (s SliceUint16) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// Zero returns the zero 0 value of the s element type.
func (s SliceUint16) Zero() interface{} {
	return 0
}
