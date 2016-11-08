package goo

import "sort"

// SliceUintptrZero is the SliceUintptr zero value.
var SliceUintptrZero = SliceUintptr(nil)

var _ Slice = SliceUintptr(nil)

var _ sort.Interface = SliceUintptr(nil)

// SliceUintptr is a slice of uintptr.
type SliceUintptr []uintptr

// Append implements Slice.
func (s SliceUintptr) Append(v ...interface{}) Slice {
	for _, v := range v {
		s = append(s, v.(uintptr))
	}

	return s
}

// AppendSlice implements Slice.
func (s SliceUintptr) AppendSlice(other Slice) Slice {
	return append(s, other.(SliceUintptr)...)
}

// Cap implements Slice.
func (s SliceUintptr) Cap() int {
	return cap(s)
}

// Copy implements Slice.
func (s SliceUintptr) Copy(other Slice) int {
	return copy(s, other.(SliceUintptr))
}

// Equals implements Equatable.
func (s SliceUintptr) Equals(other Equatable) bool {
	var t = other.(SliceUintptr)

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
func (s SliceUintptr) Get(i int) interface{} {
	return s[i]
}

// GetRange implements Slice.
func (s SliceUintptr) GetRange(i, j int) Slice {
	return s[i:j]
}

// GetRangeCap implements Slice.
func (s SliceUintptr) GetRangeCap(i, j, c int) Slice {
	return s[i:j:c]
}

// Len implements Slice.
func (s SliceUintptr) Len() int {
	return len(s)
}

// Less implements sort.Interface.
func (s SliceUintptr) Less(i, j int) bool {
	return s[i] < s[j]
}

// Make implements Slice.
func (s SliceUintptr) Make(l, c int) Slice {
	return make(SliceUintptr, l, c)
}

// NotEquals implements Equatable.
func (s SliceUintptr) NotEquals(other Equatable) bool {
	return !s.Equals(other)
}

// Set implements Slice.
func (s SliceUintptr) Set(i int, v interface{}) {
	s[i] = v.(uintptr)
}

// Swap implements sort.Interface.
func (s SliceUintptr) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
