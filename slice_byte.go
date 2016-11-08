package goo

import "sort"

// SliceByteZero is the SliceByte zero value.
var SliceByteZero = SliceByte(nil)

var _ Slice = SliceByte(nil)

var _ sort.Interface = SliceByte(nil)

// SliceByte is a slice of byte.
type SliceByte []byte

// Append implements Slice.
func (s SliceByte) Append(v ...interface{}) Slice {
	for _, v := range v {
		s = append(s, v.(byte))
	}

	return s
}

// AppendSlice implements Slice.
func (s SliceByte) AppendSlice(other Slice) Slice {
	return append(s, other.(SliceByte)...)
}

// Cap implements Slice.
func (s SliceByte) Cap() int {
	return cap(s)
}

// Copy implements Slice.
func (s SliceByte) Copy(other Slice) int {
	return copy(s, other.(SliceByte))
}

// Equals implements Equatable.
func (s SliceByte) Equals(other Equatable) bool {
	var t = other.(SliceByte)

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
func (s SliceByte) Get(i int) interface{} {
	return s[i]
}

// GetRange implements Slice.
func (s SliceByte) GetRange(i, j int) Slice {
	return s[i:j]
}

// GetRangeCap implements Slice.
func (s SliceByte) GetRangeCap(i, j, c int) Slice {
	return s[i:j:c]
}

// Len implements Slice.
func (s SliceByte) Len() int {
	return len(s)
}

// Less implements sort.Interface.
func (s SliceByte) Less(i, j int) bool {
	return s[i] < s[j]
}

// Make implements Slice.
func (s SliceByte) Make(l, c int) Slice {
	return make(SliceByte, l, c)
}

// NotEquals implements Equatable.
func (s SliceByte) NotEquals(other Equatable) bool {
	return !s.Equals(other)
}

// Set implements Slice.
func (s SliceByte) Set(i int, v interface{}) {
	s[i] = v.(byte)
}

// Swap implements sort.Interface.
func (s SliceByte) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
