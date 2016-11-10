package goo

import "sort"

// SliceInt64Zero is the SliceInt64 zero value.
var SliceInt64Zero = SliceInt64(nil)

var _ Pointer = &SliceInt64{}

var _ Slice = SliceInt64(nil)

var _ sort.Interface = SliceInt64(nil)

// SliceInt64 is a slice of int64.
type SliceInt64 []int64

// Append implements Slice.
func (s SliceInt64) Append(v ...interface{}) Slice {
	for _, v := range v {
		s = append(s, v.(int64))
	}

	return s
}

// AppendSlice implements Slice.
func (s SliceInt64) AppendSlice(other Slice) Slice {
	return append(s, other.(SliceInt64)...)
}

// Cap implements Slice.
func (s SliceInt64) Cap() int {
	return cap(s)
}

// Copy implements Slice.
func (s SliceInt64) Copy(other Slice) int {
	return copy(s, other.(SliceInt64))
}

// Dereference implements Slice.
func (s *SliceInt64) Dereference() Value {
	return *s
}

// Equals implements Slice.
func (s SliceInt64) Equals(other Equatable) bool {
	var t = other.(SliceInt64)

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
func (s SliceInt64) Get(i int) interface{} {
	return s[i]
}

// GetRange implements Slice.
func (s SliceInt64) GetRange(i, j int) Slice {
	return s[i:j]
}

// GetRangeCap implements Slice.
func (s SliceInt64) GetRangeCap(i, j, c int) Slice {
	return s[i:j:c]
}

// Len implements Slice.
func (s SliceInt64) Len() int {
	return len(s)
}

// Less implements sort.Interface.
func (s SliceInt64) Less(i, j int) bool {
	return s[i] < s[j]
}

// Make implements Slice.
func (s SliceInt64) Make(l, c int) Slice {
	return make(SliceInt64, l, c)
}

// NotEquals implements Slice.
func (s SliceInt64) NotEquals(other Equatable) bool {
	return !s.Equals(other)
}

// Reference implements Slice.
func (s SliceInt64) Reference() Pointer {
	return &s
}

// Set implements Slice.
func (s SliceInt64) Set(i int, v interface{}) {
	s[i] = v.(int64)
}

// Swap implements sort.Interface.
func (s SliceInt64) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
