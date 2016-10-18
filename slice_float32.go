package goo

import "sort"

// SliceFloat32 is a slice of float32.
type SliceFloat32 []float32

var _ Slice = SliceFloat32(nil)

var _ sort.Interface = SliceFloat32(nil)

// Append appends v to s and returns the result.
func (s SliceFloat32) Append(v ...interface{}) Slice {
	for _, v := range v {
		s = append(s, v.(float32))
	}

	return s
}

// AppendSlice appends t to s and returns the result.
func (s SliceFloat32) AppendSlice(t Slice) Slice {
	return append(s, t.(SliceFloat32)...)
}

// Cap returns the s capacity.
func (s SliceFloat32) Cap() int {
	return cap(s)
}

// Copy copies t to s.
func (s SliceFloat32) Copy(t Slice) int {
	return copy(s, t.(SliceFloat32))
}

// Equals returns whether s equals v.
func (s SliceFloat32) Equals(v interface{}) bool {
	var t = v.(SliceFloat32)
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
func (s SliceFloat32) Get(i int) interface{} {
	return s[i]
}

// GetRange returns the slice of s from indexes i to j.
func (s SliceFloat32) GetRange(i, j int) Slice {
	return s[i:j]
}

// GetRangeCap returns the slice of s from indexes i to j with capacity c.
func (s SliceFloat32) GetRangeCap(i, j, c int) Slice {
	return s[i:j:c]
}

// Len returns the s length.
func (s SliceFloat32) Len() int {
	return len(s)
}

// Less implements sort.Interface.
func (s SliceFloat32) Less(i, j int) bool {
	return s[i] < s[j]
}

// Make returns a new SliceFloat32 with length l and capacity c.
func (s SliceFloat32) Make(l, c int) Slice {
	return make(SliceFloat32, l, c)
}

// Set sets the s element at index i to v.
func (s SliceFloat32) Set(i int, v interface{}) {
	s[i] = v.(float32)
}

// Swap implements sort.Interface.
func (s SliceFloat32) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
