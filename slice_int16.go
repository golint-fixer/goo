package goo

import "sort"

// SliceInt16 is a slice of int16.
type SliceInt16 []int16

var _ sort.Interface = SliceInt16(nil)

// Append appends v to s and returns the result.
func (s SliceInt16) Append(v ...interface{}) Slice {
	for _, v := range v {
		s = append(s, v.(int16))
	}

	return s
}

// AppendSlice appends t to s and returns the result.
func (s SliceInt16) AppendSlice(t Slice) Slice {
	return append(s, t.(SliceInt16)...)
}

// Cap returns the s capacity.
func (s SliceInt16) Cap() int {
	return cap(s)
}

// Copy copies t to s.
func (s SliceInt16) Copy(t Slice) int {
	return copy(s, t.(SliceInt16))
}

// Equals returns whether s equals v.
func (s SliceInt16) Equals(v interface{}) bool {
	var t = v.(SliceInt16)
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
func (s SliceInt16) Get(i int) interface{} {
	return s[i]
}

// Len returns the s length.
func (s SliceInt16) Len() int {
	return len(s)
}

// Less implements sort.Interface.
func (s SliceInt16) Less(i, j int) bool {
	return s[i] < s[j]
}

// Make returns a new SliceInt16 with length l and capacity c.
func (s SliceInt16) Make(l, c int) Slice {
	return make(SliceInt16, l, c)
}

// Set sets the s element at index i to v.
func (s SliceInt16) Set(i int, v interface{}) {
	s[i] = v.(int16)
}

// Slice returns the slice of s from indexes i to j.
func (s SliceInt16) Slice(i, j int) Slice {
	return s[i:j]
}

// SliceCap returns the slice of s from indexes i to j with capacity c.
func (s SliceInt16) SliceCap(i, j, c int) Slice {
	return s[i:j:c]
}

// Swap implements sort.Interface.
func (s SliceInt16) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// Zero returns the zero 0 value of the s element type.
func (s SliceInt16) Zero() interface{} {
	return 0
}
