package goo

import "sort"

// SliceUint64Zero is the SliceUint64 zero value.
var SliceUint64Zero = SliceUint64(nil)

var _ Pointer = &SliceUint64{}

var _ Slice = SliceUint64(nil)

var _ sort.Interface = SliceUint64(nil)

// SliceUint64 is a slice of uint64.
type SliceUint64 []uint64

// Append implements Slice.
func (s SliceUint64) Append(v ...interface{}) Slice {
	for _, v := range v {
		s = append(s, v.(uint64))
	}

	return s
}

// AppendSlice implements Slice.
func (s SliceUint64) AppendSlice(other Slice) Slice {
	return append(s, other.(SliceUint64)...)
}

// Cap implements Slice.
func (s SliceUint64) Cap() int {
	return cap(s)
}

// Copy implements Slice.
func (s SliceUint64) Copy(other Slice) int {
	return copy(s, other.(SliceUint64))
}

// Dereference implements Slice.
func (s *SliceUint64) Dereference() Value {
	return *s
}

// Equals implements Slice.
func (s SliceUint64) Equals(other Equatable) bool {
	var t = other.(SliceUint64)

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
func (s SliceUint64) Get(i int) interface{} {
	return s[i]
}

// GetRange implements Slice.
func (s SliceUint64) GetRange(i, j int) Slice {
	return s[i:j]
}

// GetRangeCap implements Slice.
func (s SliceUint64) GetRangeCap(i, j, c int) Slice {
	return s[i:j:c]
}

// Len implements Slice.
func (s SliceUint64) Len() int {
	return len(s)
}

// Less implements sort.Interface.
func (s SliceUint64) Less(i, j int) bool {
	return s[i] < s[j]
}

// Make implements Slice.
func (s SliceUint64) Make(l, c int) Slice {
	return make(SliceUint64, l, c)
}

// NotEquals implements Slice.
func (s SliceUint64) NotEquals(other Equatable) bool {
	return !s.Equals(other)
}

// Reference implements Slice.
func (s SliceUint64) Reference() Pointer {
	return &s
}

// Set implements Slice.
func (s SliceUint64) Set(i int, v interface{}) {
	s[i] = v.(uint64)
}

// Swap implements sort.Interface.
func (s SliceUint64) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
