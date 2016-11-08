package goo

import "sort"

// SliceStringZero is the SliceString zero value.
var SliceStringZero = SliceString(nil)

var _ Slice = SliceString(nil)

var _ sort.Interface = SliceString(nil)

// SliceString is a slice of string.
type SliceString []string

// Append implements Slice.
func (s SliceString) Append(v ...interface{}) Slice {
	for _, v := range v {
		s = append(s, v.(string))
	}

	return s
}

// AppendSlice implements Slice.
func (s SliceString) AppendSlice(other Slice) Slice {
	return append(s, other.(SliceString)...)
}

// Cap implements Slice.
func (s SliceString) Cap() int {
	return cap(s)
}

// Copy implements Slice.
func (s SliceString) Copy(other Slice) int {
	return copy(s, other.(SliceString))
}

// Equals implements Equatable.
func (s SliceString) Equals(other Equatable) bool {
	var t = other.(SliceString)

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
func (s SliceString) Get(i int) interface{} {
	return s[i]
}

// GetRange implements Slice.
func (s SliceString) GetRange(i, j int) Slice {
	return s[i:j]
}

// GetRangeCap implements Slice.
func (s SliceString) GetRangeCap(i, j, c int) Slice {
	return s[i:j:c]
}

// Len implements Slice.
func (s SliceString) Len() int {
	return len(s)
}

// Less implements sort.Interface.
func (s SliceString) Less(i, j int) bool {
	return s[i] < s[j]
}

// Make implements Slice.
func (s SliceString) Make(l, c int) Slice {
	return make(SliceString, l, c)
}

// NotEquals implements Equatable.
func (s SliceString) NotEquals(other Equatable) bool {
	return !s.Equals(other)
}

// Set implements Slice.
func (s SliceString) Set(i int, v interface{}) {
	s[i] = v.(string)
}

// Swap implements sort.Interface.
func (s SliceString) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
