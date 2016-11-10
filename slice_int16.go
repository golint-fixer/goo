package goo

import "sort"

// SliceInt16Zero is the SliceInt16 zero value.
var SliceInt16Zero = SliceInt16(nil)

var _ Slice = SliceInt16(nil)

var _ sort.Interface = SliceInt16(nil)

// SliceInt16 is a slice of int16.
type SliceInt16 []int16

// Append implements Slice.
func (s SliceInt16) Append(v ...interface{}) Slice {
	for _, v := range v {
		s = append(s, v.(int16))
	}

	return s
}

// AppendSlice implements Slice.
func (s SliceInt16) AppendSlice(other Slice) Slice {
	return append(s, other.(SliceInt16)...)
}

// Cap implements Slice.
func (s SliceInt16) Cap() int {
	return cap(s)
}

// Copy implements Slice.
func (s SliceInt16) Copy(other Slice) int {
	return copy(s, other.(SliceInt16))
}

// Dereference implements Slice.
func (s *SliceInt16) Dereference() Value {
	return *s
}

// Equals implements Slice.
func (s SliceInt16) Equals(other Equatable) bool {
	var t = other.(SliceInt16)

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
func (s SliceInt16) Get(i int) interface{} {
	return s[i]
}

// GetRange implements Slice.
func (s SliceInt16) GetRange(i, j int) Slice {
	return s[i:j]
}

// GetRangeCap implements Slice.
func (s SliceInt16) GetRangeCap(i, j, c int) Slice {
	return s[i:j:c]
}

// Len implements Slice.
func (s SliceInt16) Len() int {
	return len(s)
}

// Less implements sort.Interface.
func (s SliceInt16) Less(i, j int) bool {
	return s[i] < s[j]
}

// Make implements Slice.
func (s SliceInt16) Make(l, c int) Slice {
	return make(SliceInt16, l, c)
}

// NotEquals implements Slice.
func (s SliceInt16) NotEquals(other Equatable) bool {
	return !s.Equals(other)
}

// Reference implements Slice.
func (s SliceInt16) Reference() Pointer {
	return &s
}

// Set implements Slice.
func (s SliceInt16) Set(i int, v interface{}) {
	s[i] = v.(int16)
}

// Swap implements sort.Interface.
func (s SliceInt16) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
