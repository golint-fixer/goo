package goo

import "sort"

// SliceInt32 is a slice of int32.
type SliceInt32 []int32

var _ sort.Interface = SliceInt32(nil)

// Append appends v to s and returns the result.
func (s SliceInt32) Append(v ...interface{}) Slice {
	for _, v := range v {
		s = append(s, v.(int32))
	}

	return s
}

// AppendSlice appends t to s and returns the result.
func (s SliceInt32) AppendSlice(t Slice) Slice {
	return append(s, t.(SliceInt32)...)
}

// Cap returns the s capacity.
func (s SliceInt32) Cap() int {
	return cap(s)
}

// Copy copies t to s.
func (s SliceInt32) Copy(t Slice) int {
	return copy(s, t.(SliceInt32))
}

// Equals returns whether s equals v.
func (s SliceInt32) Equals(v interface{}) bool {
	var t = v.(SliceInt32)
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
func (s SliceInt32) Get(i int) interface{} {
	return s[i]
}

// Len returns the s length.
func (s SliceInt32) Len() int {
	return len(s)
}

// Less implements sort.Interface.
func (s SliceInt32) Less(i, j int) bool {
	return s[i] < s[j]
}

// Make returns a new SliceInt32 with length l and capacity c.
func (s SliceInt32) Make(l, c int) Slice {
	return make(SliceInt32, l, c)
}

// Set sets the s element at index i to v.
func (s SliceInt32) Set(i int, v interface{}) {
	s[i] = v.(int32)
}

// Slice returns the slice of s from indexes i to j.
func (s SliceInt32) Slice(i, j int) Slice {
	return s[i:j]
}

// SliceCap returns the slice of s from indexes i to j with capacity c.
func (s SliceInt32) SliceCap(i, j, c int) Slice {
	return s[i:j:c]
}

// Swap implements sort.Interface.
func (s SliceInt32) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// Zero returns the zero 0 value of the s element type.
func (s SliceInt32) Zero() interface{} {
	return 0
}
