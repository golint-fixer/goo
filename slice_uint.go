package goo

import "sort"

// SliceUintZero is the SliceUint zero value.
var SliceUintZero = SliceUint(nil)

var _ Slice = SliceUint(nil)

var _ sort.Interface = SliceUint(nil)

// SliceUint is a slice of uint.
type SliceUint []uint

// Append implements Slice.
func (s SliceUint) Append(v ...interface{}) Slice {
	for _, v := range v {
		s = append(s, v.(uint))
	}

	return s
}

// AppendSlice implements Slice.
func (s SliceUint) AppendSlice(other Slice) Slice {
	return append(s, other.(SliceUint)...)
}

// Cap implements Slice.
func (s SliceUint) Cap() int {
	return cap(s)
}

// Copy implements Slice.
func (s SliceUint) Copy(other Slice) int {
	return copy(s, other.(SliceUint))
}

// Equals implements Equatable.
func (s SliceUint) Equals(other Equatable) bool {
	var t = other.(SliceUint)

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
func (s SliceUint) Get(i int) interface{} {
	return s[i]
}

// GetRange implements Slice.
func (s SliceUint) GetRange(i, j int) Slice {
	return s[i:j]
}

// GetRangeCap implements Slice.
func (s SliceUint) GetRangeCap(i, j, c int) Slice {
	return s[i:j:c]
}

// Len implements Slice.
func (s SliceUint) Len() int {
	return len(s)
}

// Less implements sort.Interface.
func (s SliceUint) Less(i, j int) bool {
	return s[i] < s[j]
}

// Make implements Slice.
func (s SliceUint) Make(l, c int) Slice {
	return make(SliceUint, l, c)
}

// NotEquals implements Equatable.
func (s SliceUint) NotEquals(other Equatable) bool {
	return !s.Equals(other)
}

// Set implements Slice.
func (s SliceUint) Set(i int, v interface{}) {
	s[i] = v.(uint)
}

// Swap implements sort.Interface.
func (s SliceUint) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
