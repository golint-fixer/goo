package goo

import "sort"

// SliceIntZero is the SliceInt zero value.
var SliceIntZero = SliceInt(nil)

var _ Slice = SliceInt(nil)

var _ sort.Interface = SliceInt(nil)

// SliceInt is a slice of int.
type SliceInt []int

// Append implements Slice.
func (s SliceInt) Append(v ...interface{}) Slice {
	for _, v := range v {
		s = append(s, v.(int))
	}

	return s
}

// AppendSlice implements Slice.
func (s SliceInt) AppendSlice(other Slice) Slice {
	return append(s, other.(SliceInt)...)
}

// Cap implements Slice.
func (s SliceInt) Cap() int {
	return cap(s)
}

// Copy implements Slice.
func (s SliceInt) Copy(other Slice) int {
	return copy(s, other.(SliceInt))
}

// Equals implements Slice.
func (s SliceInt) Equals(other Equatable) bool {
	var t = other.(SliceInt)

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
func (s SliceInt) Get(i int) interface{} {
	return s[i]
}

// GetRange implements Slice.
func (s SliceInt) GetRange(i, j int) Slice {
	return s[i:j]
}

// GetRangeCap implements Slice.
func (s SliceInt) GetRangeCap(i, j, c int) Slice {
	return s[i:j:c]
}

// Len implements Slice.
func (s SliceInt) Len() int {
	return len(s)
}

// Less implements sort.Interface.
func (s SliceInt) Less(i, j int) bool {
	return s[i] < s[j]
}

// Make implements Slice.
func (s SliceInt) Make(l, c int) Slice {
	return make(SliceInt, l, c)
}

// NotEquals implements Slice.
func (s SliceInt) NotEquals(other Equatable) bool {
	return !s.Equals(other)
}

// Set implements Slice.
func (s SliceInt) Set(i int, v interface{}) {
	s[i] = v.(int)
}

// Swap implements sort.Interface.
func (s SliceInt) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
