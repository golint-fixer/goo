package goo

import "sort"

// SliceFloat64Zero is the SliceFloat64 zero value.
var SliceFloat64Zero = SliceFloat64(nil)

var _ Slice = SliceFloat64(nil)

var _ sort.Interface = SliceFloat64(nil)

// SliceFloat64 is a slice of float64.
type SliceFloat64 []float64

// Append implements Slice.
func (s SliceFloat64) Append(v ...interface{}) Slice {
	for _, v := range v {
		s = append(s, v.(float64))
	}

	return s
}

// AppendSlice implements Slice.
func (s SliceFloat64) AppendSlice(other Slice) Slice {
	return append(s, other.(SliceFloat64)...)
}

// Cap implements Slice.
func (s SliceFloat64) Cap() int {
	return cap(s)
}

// Copy implements Slice.
func (s SliceFloat64) Copy(other Slice) int {
	return copy(s, other.(SliceFloat64))
}

// Equals implements Equatable.
func (s SliceFloat64) Equals(other Equatable) bool {
	var t = other.(SliceFloat64)

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
func (s SliceFloat64) Get(i int) interface{} {
	return s[i]
}

// GetRange implements Slice.
func (s SliceFloat64) GetRange(i, j int) Slice {
	return s[i:j]
}

// GetRangeCap implements Slice.
func (s SliceFloat64) GetRangeCap(i, j, c int) Slice {
	return s[i:j:c]
}

// Len implements Slice.
func (s SliceFloat64) Len() int {
	return len(s)
}

// Less implements sort.Interface.
func (s SliceFloat64) Less(i, j int) bool {
	return s[i] < s[j]
}

// Make implements Slice.
func (s SliceFloat64) Make(l, c int) Slice {
	return make(SliceFloat64, l, c)
}

// NotEquals implements Equatable.
func (s SliceFloat64) NotEquals(other Equatable) bool {
	return !s.Equals(other)
}

// Set implements Slice.
func (s SliceFloat64) Set(i int, v interface{}) {
	s[i] = v.(float64)
}

// Swap implements sort.Interface.
func (s SliceFloat64) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
