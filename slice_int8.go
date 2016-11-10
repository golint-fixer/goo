package goo

import "sort"

// SliceInt8Zero is the SliceInt8 zero value.
var SliceInt8Zero = SliceInt8(nil)

var _ Slice = SliceInt8(nil)

var _ sort.Interface = SliceInt8(nil)

// SliceInt8 is a slice of int8.
type SliceInt8 []int8

// Append implements Slice.
func (s SliceInt8) Append(v ...interface{}) Slice {
	for _, v := range v {
		s = append(s, v.(int8))
	}

	return s
}

// AppendSlice implements Slice.
func (s SliceInt8) AppendSlice(other Slice) Slice {
	return append(s, other.(SliceInt8)...)
}

// Cap implements Slice.
func (s SliceInt8) Cap() int {
	return cap(s)
}

// Copy implements Slice.
func (s SliceInt8) Copy(other Slice) int {
	return copy(s, other.(SliceInt8))
}

// Dereference implements Slice.
func (s *SliceInt8) Dereference() Value {
	return *s
}

// Equals implements Slice.
func (s SliceInt8) Equals(other Equatable) bool {
	var t = other.(SliceInt8)

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
func (s SliceInt8) Get(i int) interface{} {
	return s[i]
}

// GetRange implements Slice.
func (s SliceInt8) GetRange(i, j int) Slice {
	return s[i:j]
}

// GetRangeCap implements Slice.
func (s SliceInt8) GetRangeCap(i, j, c int) Slice {
	return s[i:j:c]
}

// Len implements Slice.
func (s SliceInt8) Len() int {
	return len(s)
}

// Less implements sort.Interface.
func (s SliceInt8) Less(i, j int) bool {
	return s[i] < s[j]
}

// Make implements Slice.
func (s SliceInt8) Make(l, c int) Slice {
	return make(SliceInt8, l, c)
}

// NotEquals implements Slice.
func (s SliceInt8) NotEquals(other Equatable) bool {
	return !s.Equals(other)
}

// Reference implements Slice.
func (s SliceInt8) Reference() Pointer {
	return &s
}

// Set implements Slice.
func (s SliceInt8) Set(i int, v interface{}) {
	s[i] = v.(int8)
}

// Swap implements sort.Interface.
func (s SliceInt8) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
