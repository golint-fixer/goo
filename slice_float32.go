package goo

import "sort"

// SliceFloat32Zero is the SliceFloat32 zero value.
var SliceFloat32Zero = SliceFloat32(nil)

var _ Slice = SliceFloat32(nil)

var _ sort.Interface = SliceFloat32(nil)

// SliceFloat32 is a slice of float32.
type SliceFloat32 []float32

// Append implements Slice.
func (s SliceFloat32) Append(v ...interface{}) Slice {
	for _, v := range v {
		s = append(s, v.(float32))
	}

	return s
}

// AppendSlice implements Slice.
func (s SliceFloat32) AppendSlice(other Slice) Slice {
	return append(s, other.(SliceFloat32)...)
}

// Cap implements Slice.
func (s SliceFloat32) Cap() int {
	return cap(s)
}

// Copy implements Slice.
func (s SliceFloat32) Copy(other Slice) int {
	return copy(s, other.(SliceFloat32))
}

// Equals implements Slice.
func (s SliceFloat32) Equals(other Equatable) bool {
	var t = other.(SliceFloat32)

	if len(t) != len(s) {
		return false
	}

	for i := range s {
		if t[i] != s[i] {
			return false
		}
	}

	return true
}

// Get implements Slice.
func (s SliceFloat32) Get(i int) interface{} {
	return s[i]
}

// GetRange implements Slice.
func (s SliceFloat32) GetRange(i, j int) Slice {
	return s[i:j]
}

// GetRangeCap implements Slice.
func (s SliceFloat32) GetRangeCap(i, j, c int) Slice {
	return s[i:j:c]
}

// Len implements Slice.
func (s SliceFloat32) Len() int {
	return len(s)
}

// Less implements sort.Interface.
func (s SliceFloat32) Less(i, j int) bool {
	return s[i] < s[j]
}

// Make implements Slice.
func (s SliceFloat32) Make(l, c int) Slice {
	return make(SliceFloat32, l, c)
}

// NotEquals implements Slice.
func (s SliceFloat32) NotEquals(other Equatable) bool {
	return !s.Equals(other)
}

// Set implements Slice.
func (s SliceFloat32) Set(i int, v interface{}) {
	s[i] = v.(float32)
}

// Swap implements sort.Interface.
func (s SliceFloat32) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
