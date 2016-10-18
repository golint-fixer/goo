package goo

import "sort"

// SliceInt8 is a slice of int8.
type SliceInt8 []int8

var _ Slice = SliceInt8(nil)

var _ sort.Interface = SliceInt8(nil)

// Append appends v to s and returns the result.
func (s SliceInt8) Append(v ...interface{}) Slice {
	for _, v := range v {
		s = append(s, v.(int8))
	}

	return s
}

// AppendSlice appends t to s and returns the result.
func (s SliceInt8) AppendSlice(t Slice) Slice {
	return append(s, t.(SliceInt8)...)
}

// Cap returns the s capacity.
func (s SliceInt8) Cap() int {
	return cap(s)
}

// Copy copies t to s.
func (s SliceInt8) Copy(t Slice) int {
	return copy(s, t.(SliceInt8))
}

// Equals returns whether s equals v.
func (s SliceInt8) Equals(v interface{}) bool {
	var t = v.(SliceInt8)
	var l = len(s)

	if len(t) != l {
		return false
	}

	if l == 0 {
		return true
	}

	for i := range s {
		if t[i] != s[i] {
			return false
		}
	}

	return true
}

// Get returns the s element at index i.
func (s SliceInt8) Get(i int) interface{} {
	return s[i]
}

// GetRange returns the slice of s from indexes i to j.
func (s SliceInt8) GetRange(i, j int) Slice {
	return s[i:j]
}

// GetRangeCap returns the slice of s from indexes i to j with capacity c.
func (s SliceInt8) GetRangeCap(i, j, c int) Slice {
	return s[i:j:c]
}

// Len returns the s length.
func (s SliceInt8) Len() int {
	return len(s)
}

// Less implements sort.Interface.
func (s SliceInt8) Less(i, j int) bool {
	return s[i] < s[j]
}

// Make returns a new SliceInt8 with length l and capacity c.
func (s SliceInt8) Make(l, c int) Slice {
	return make(SliceInt8, l, c)
}

// Set sets the s element at index i to v.
func (s SliceInt8) Set(i int, v interface{}) {
	s[i] = v.(int8)
}

// Swap implements sort.Interface.
func (s SliceInt8) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
