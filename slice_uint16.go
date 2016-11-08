package goo

import "sort"

// SliceUint16Zero is the SliceUint16 zero value.
var SliceUint16Zero = SliceUint16(nil)

var _ Slice = SliceUint16(nil)

var _ sort.Interface = SliceUint16(nil)

// SliceUint16 is a slice of uint16.
type SliceUint16 []uint16

// Append implements Slice.
func (s SliceUint16) Append(v ...interface{}) Slice {
	for _, v := range v {
		s = append(s, v.(uint16))
	}

	return s
}

// AppendSlice implements Slice.
func (s SliceUint16) AppendSlice(other Slice) Slice {
	return append(s, other.(SliceUint16)...)
}

// Cap implements Slice.
func (s SliceUint16) Cap() int {
	return cap(s)
}

// Copy implements Slice.
func (s SliceUint16) Copy(other Slice) int {
	return copy(s, other.(SliceUint16))
}

// Equals implements Equatable.
func (s SliceUint16) Equals(other Equatable) bool {
	var t = other.(SliceUint16)

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
func (s SliceUint16) Get(i int) interface{} {
	return s[i]
}

// GetRange implements Slice.
func (s SliceUint16) GetRange(i, j int) Slice {
	return s[i:j]
}

// GetRangeCap implements Slice.
func (s SliceUint16) GetRangeCap(i, j, c int) Slice {
	return s[i:j:c]
}

// Len implements Slice.
func (s SliceUint16) Len() int {
	return len(s)
}

// Less implements sort.Interface.
func (s SliceUint16) Less(i, j int) bool {
	return s[i] < s[j]
}

// Make implements Slice.
func (s SliceUint16) Make(l, c int) Slice {
	return make(SliceUint16, l, c)
}

// NotEquals implements Equatable.
func (s SliceUint16) NotEquals(other Equatable) bool {
	return !s.Equals(other)
}

// Set implements Slice.
func (s SliceUint16) Set(i int, v interface{}) {
	s[i] = v.(uint16)
}

// Swap implements sort.Interface.
func (s SliceUint16) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
