package goo

import "sort"

// SliceInt32Zero is the SliceInt32 zero value.
var SliceInt32Zero = SliceInt32(nil)

var _ Slice = SliceInt32(nil)

var _ sort.Interface = SliceInt32(nil)

// SliceInt32 is a slice of int32.
type SliceInt32 []int32

// Append implements Slice.
func (s SliceInt32) Append(v ...interface{}) Slice {
	for _, v := range v {
		s = append(s, v.(int32))
	}

	return s
}

// AppendSlice implements Slice.
func (s SliceInt32) AppendSlice(other Slice) Slice {
	return append(s, other.(SliceInt32)...)
}

// Cap implements Slice.
func (s SliceInt32) Cap() int {
	return cap(s)
}

// Copy implements Slice.
func (s SliceInt32) Copy(other Slice) int {
	return copy(s, other.(SliceInt32))
}

// Equals implements Equatable.
func (s SliceInt32) Equals(other Equatable) bool {
	var t = other.(SliceInt32)

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
func (s SliceInt32) Get(i int) interface{} {
	return s[i]
}

// GetRange implements Slice.
func (s SliceInt32) GetRange(i, j int) Slice {
	return s[i:j]
}

// GetRangeCap implements Slice.
func (s SliceInt32) GetRangeCap(i, j, c int) Slice {
	return s[i:j:c]
}

// Len implements Slice.
func (s SliceInt32) Len() int {
	return len(s)
}

// Less implements sort.Interface.
func (s SliceInt32) Less(i, j int) bool {
	return s[i] < s[j]
}

// Make implements Slice.
func (s SliceInt32) Make(l, c int) Slice {
	return make(SliceInt32, l, c)
}

// NotEquals implements Equatable.
func (s SliceInt32) NotEquals(other Equatable) bool {
	return !s.Equals(other)
}

// Set implements Slice.
func (s SliceInt32) Set(i int, v interface{}) {
	s[i] = v.(int32)
}

// Swap implements sort.Interface.
func (s SliceInt32) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
