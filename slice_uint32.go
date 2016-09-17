package goo

import "sort"

// SliceUint32 is a slice of uint32.
type SliceUint32 []uint32

var _ sort.Interface = SliceUint32(nil)

// Append appends v to s and returns the result.
func (s SliceUint32) Append(v ...interface{}) Slice {
	for _, v := range v {
		s = append(s, v.(uint32))
	}

	return s
}

// AppendSlice appends t to s and returns the result.
func (s SliceUint32) AppendSlice(t Slice) Slice {
	return append(s, t.(SliceUint32)...)
}

// Cap returns the s capacity.
func (s SliceUint32) Cap() int {
	return cap(s)
}

// Copy copies t to s.
func (s SliceUint32) Copy(t Slice) int {
	return copy(s, t.(SliceUint32))
}

// Equals returns whether s equals v.
func (s SliceUint32) Equals(v interface{}) bool {
	var t = v.(SliceUint32)
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
func (s SliceUint32) Get(i int) interface{} {
	return s[i]
}

// Len returns the s length.
func (s SliceUint32) Len() int {
	return len(s)
}

// Less implements sort.Interface.
func (s SliceUint32) Less(i, j int) bool {
	return s[i] < s[j]
}

// Make returns a new SliceUint32 with length l and capacity c.
func (s SliceUint32) Make(l, c int) Slice {
	return make(SliceUint32, l, c)
}

// Set sets the s element at index i to v.
func (s SliceUint32) Set(i int, v interface{}) {
	s[i] = v.(uint32)
}

// Slice returns the slice of s from indexes i to j.
func (s SliceUint32) Slice(i, j int) Slice {
	return s[i:j]
}

// SliceCap returns the slice of s from indexes i to j with capacity c.
func (s SliceUint32) SliceCap(i, j, c int) Slice {
	return s[i:j:c]
}

// Swap implements sort.Interface.
func (s SliceUint32) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// Zero returns the zero 0 value of the s element type.
func (s SliceUint32) Zero() interface{} {
	return 0
}
