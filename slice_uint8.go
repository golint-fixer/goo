package goo

import "sort"

// SliceUint8 is a slice of uint8.
type SliceUint8 []uint8

var _ Slice = SliceUint8(nil)

var _ sort.Interface = SliceUint8(nil)

// Append appends v to s and returns the result.
func (s SliceUint8) Append(v ...interface{}) Slice {
	for _, v := range v {
		s = append(s, v.(uint8))
	}

	return s
}

// AppendSlice appends t to s and returns the result.
func (s SliceUint8) AppendSlice(t Slice) Slice {
	return append(s, t.(SliceUint8)...)
}

// Cap returns the s capacity.
func (s SliceUint8) Cap() int {
	return cap(s)
}

// Copy copies t to s.
func (s SliceUint8) Copy(t Slice) int {
	return copy(s, t.(SliceUint8))
}

// Equals returns whether s equals v.
func (s SliceUint8) Equals(v interface{}) bool {
	var t = v.(SliceUint8)
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
func (s SliceUint8) Get(i int) interface{} {
	return s[i]
}

// Len returns the s length.
func (s SliceUint8) Len() int {
	return len(s)
}

// Less implements sort.Interface.
func (s SliceUint8) Less(i, j int) bool {
	return s[i] < s[j]
}

// Make returns a new SliceUint8 with length l and capacity c.
func (s SliceUint8) Make(l, c int) Slice {
	return make(SliceUint8, l, c)
}

// Set sets the s element at index i to v.
func (s SliceUint8) Set(i int, v interface{}) {
	s[i] = v.(uint8)
}

// Slice returns the slice of s from indexes i to j.
func (s SliceUint8) Slice(i, j int) Slice {
	return s[i:j]
}

// SliceCap returns the slice of s from indexes i to j with capacity c.
func (s SliceUint8) SliceCap(i, j, c int) Slice {
	return s[i:j:c]
}

// Swap implements sort.Interface.
func (s SliceUint8) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
