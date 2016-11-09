package goo

import "sort"

// SliceUint8Zero is the SliceUint8 zero value.
var SliceUint8Zero = SliceUint8(nil)

var _ Slice = SliceUint8(nil)

var _ sort.Interface = SliceUint8(nil)

// SliceUint8 is a slice of uint8.
type SliceUint8 []uint8

// Append implements Slice.
func (s SliceUint8) Append(v ...interface{}) Slice {
	for _, v := range v {
		s = append(s, v.(uint8))
	}

	return s
}

// AppendSlice implements Slice.
func (s SliceUint8) AppendSlice(other Slice) Slice {
	return append(s, other.(SliceUint8)...)
}

// Cap implements Slice.
func (s SliceUint8) Cap() int {
	return cap(s)
}

// Copy implements Slice.
func (s SliceUint8) Copy(other Slice) int {
	return copy(s, other.(SliceUint8))
}

// Equals implements Slice.
func (s SliceUint8) Equals(other Equatable) bool {
	var t = other.(SliceUint8)

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
func (s SliceUint8) Get(i int) interface{} {
	return s[i]
}

// GetRange implements Slice.
func (s SliceUint8) GetRange(i, j int) Slice {
	return s[i:j]
}

// GetRangeCap implements Slice.
func (s SliceUint8) GetRangeCap(i, j, c int) Slice {
	return s[i:j:c]
}

// Len implements Slice.
func (s SliceUint8) Len() int {
	return len(s)
}

// Less implements sort.Interface.
func (s SliceUint8) Less(i, j int) bool {
	return s[i] < s[j]
}

// Make implements Slice.
func (s SliceUint8) Make(l, c int) Slice {
	return make(SliceUint8, l, c)
}

// NotEquals implements Slice.
func (s SliceUint8) NotEquals(other Equatable) bool {
	return !s.Equals(other)
}

// Set implements Slice.
func (s SliceUint8) Set(i int, v interface{}) {
	s[i] = v.(uint8)
}

// Swap implements sort.Interface.
func (s SliceUint8) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
