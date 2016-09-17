package goo

import "sort"

// SliceUintptr is a slice of uintptr.
type SliceUintptr []uintptr

var _ sort.Interface = SliceUintptr(nil)

// Append appends v to s and returns the result.
func (s SliceUintptr) Append(v ...interface{}) Slice {
	for _, v := range v {
		s = append(s, v.(uintptr))
	}

	return s
}

// AppendSlice appends t to s and returns the result.
func (s SliceUintptr) AppendSlice(t Slice) Slice {
	return append(s, t.(SliceUintptr)...)
}

// Cap returns the s capacity.
func (s SliceUintptr) Cap() int {
	return cap(s)
}

// Copy copies t to s.
func (s SliceUintptr) Copy(t Slice) int {
	return copy(s, t.(SliceUintptr))
}

// Equals returns whether s equals v.
func (s SliceUintptr) Equals(v interface{}) bool {
	var t = v.(SliceUintptr)
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
func (s SliceUintptr) Get(i int) interface{} {
	return s[i]
}

// Len returns the s length.
func (s SliceUintptr) Len() int {
	return len(s)
}

// Less implements sort.Interface.
func (s SliceUintptr) Less(i, j int) bool {
	return s[i] < s[j]
}

// Make returns a new SliceUintptr with length l and capacity c.
func (s SliceUintptr) Make(l, c int) Slice {
	return make(SliceUintptr, l, c)
}

// Set sets the s element at index i to v.
func (s SliceUintptr) Set(i int, v interface{}) {
	s[i] = v.(uintptr)
}

// Slice returns the slice of s from indexes i to j.
func (s SliceUintptr) Slice(i, j int) Slice {
	return s[i:j]
}

// SliceCap returns the slice of s from indexes i to j with capacity c.
func (s SliceUintptr) SliceCap(i, j, c int) Slice {
	return s[i:j:c]
}

// Swap implements sort.Interface.
func (s SliceUintptr) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// Zero returns the zero 0 value of the s element type.
func (s SliceUintptr) Zero() interface{} {
	return 0
}
