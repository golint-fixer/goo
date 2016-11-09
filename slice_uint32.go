package goo

import "sort"

// SliceUint32Zero is the SliceUint32 zero value.
var SliceUint32Zero = SliceUint32(nil)

var _ Slice = SliceUint32(nil)

var _ sort.Interface = SliceUint32(nil)

// SliceUint32 is a slice of uint32.
type SliceUint32 []uint32

// Append implements Slice.
func (s SliceUint32) Append(v ...interface{}) Slice {
	for _, v := range v {
		s = append(s, v.(uint32))
	}

	return s
}

// AppendSlice implements Slice.
func (s SliceUint32) AppendSlice(other Slice) Slice {
	return append(s, other.(SliceUint32)...)
}

// Cap implements Slice.
func (s SliceUint32) Cap() int {
	return cap(s)
}

// Copy implements Slice.
func (s SliceUint32) Copy(other Slice) int {
	return copy(s, other.(SliceUint32))
}

// Equals implements Slice.
func (s SliceUint32) Equals(other Equatable) bool {
	var t = other.(SliceUint32)

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
func (s SliceUint32) Get(i int) interface{} {
	return s[i]
}

// GetRange implements Slice.
func (s SliceUint32) GetRange(i, j int) Slice {
	return s[i:j]
}

// GetRangeCap implements Slice.
func (s SliceUint32) GetRangeCap(i, j, c int) Slice {
	return s[i:j:c]
}

// Len implements Slice.
func (s SliceUint32) Len() int {
	return len(s)
}

// Less implements sort.Interface.
func (s SliceUint32) Less(i, j int) bool {
	return s[i] < s[j]
}

// Make implements Slice.
func (s SliceUint32) Make(l, c int) Slice {
	return make(SliceUint32, l, c)
}

// NotEquals implements Slice.
func (s SliceUint32) NotEquals(other Equatable) bool {
	return !s.Equals(other)
}

// Set implements Slice.
func (s SliceUint32) Set(i int, v interface{}) {
	s[i] = v.(uint32)
}

// Swap implements sort.Interface.
func (s SliceUint32) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
