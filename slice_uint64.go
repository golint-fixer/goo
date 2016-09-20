package goo

import "sort"

// SliceUint64 is a slice of uint64.
type SliceUint64 []uint64

var _ Slice = SliceUint64(nil)

var _ sort.Interface = SliceUint64(nil)

// Append appends v to s and returns the result.
func (s SliceUint64) Append(v ...interface{}) Slice {
	for _, v := range v {
		s = append(s, v.(uint64))
	}

	return s
}

// AppendSlice appends t to s and returns the result.
func (s SliceUint64) AppendSlice(t Slice) Slice {
	return append(s, t.(SliceUint64)...)
}

// Cap returns the s capacity.
func (s SliceUint64) Cap() int {
	return cap(s)
}

// Copy copies t to s.
func (s SliceUint64) Copy(t Slice) int {
	return copy(s, t.(SliceUint64))
}

// Equals returns whether s equals v.
func (s SliceUint64) Equals(v interface{}) bool {
	var t = v.(SliceUint64)
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
func (s SliceUint64) Get(i int) interface{} {
	return s[i]
}

// Len returns the s length.
func (s SliceUint64) Len() int {
	return len(s)
}

// Less implements sort.Interface.
func (s SliceUint64) Less(i, j int) bool {
	return s[i] < s[j]
}

// Make returns a new SliceUint64 with length l and capacity c.
func (s SliceUint64) Make(l, c int) Slice {
	return make(SliceUint64, l, c)
}

// Set sets the s element at index i to v.
func (s SliceUint64) Set(i int, v interface{}) {
	s[i] = v.(uint64)
}

// Slice returns the slice of s from indexes i to j.
func (s SliceUint64) Slice(i, j int) Slice {
	return s[i:j]
}

// SliceCap returns the slice of s from indexes i to j with capacity c.
func (s SliceUint64) SliceCap(i, j, c int) Slice {
	return s[i:j:c]
}

// Swap implements sort.Interface.
func (s SliceUint64) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// Zero returns the zero 0 value of the s element type.
func (s SliceUint64) Zero() interface{} {
	return 0
}
