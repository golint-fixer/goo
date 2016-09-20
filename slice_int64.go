package goo

import "sort"

// SliceInt64 is a slice of int64.
type SliceInt64 []int64

var _ Slice = SliceInt64(nil)

var _ sort.Interface = SliceInt64(nil)

// Append appends v to s and returns the result.
func (s SliceInt64) Append(v ...interface{}) Slice {
	for _, v := range v {
		s = append(s, v.(int64))
	}

	return s
}

// AppendSlice appends t to s and returns the result.
func (s SliceInt64) AppendSlice(t Slice) Slice {
	return append(s, t.(SliceInt64)...)
}

// Cap returns the s capacity.
func (s SliceInt64) Cap() int {
	return cap(s)
}

// Copy copies t to s.
func (s SliceInt64) Copy(t Slice) int {
	return copy(s, t.(SliceInt64))
}

// Equals returns whether s equals v.
func (s SliceInt64) Equals(v interface{}) bool {
	var t = v.(SliceInt64)
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
func (s SliceInt64) Get(i int) interface{} {
	return s[i]
}

// Len returns the s length.
func (s SliceInt64) Len() int {
	return len(s)
}

// Less implements sort.Interface.
func (s SliceInt64) Less(i, j int) bool {
	return s[i] < s[j]
}

// Make returns a new SliceInt64 with length l and capacity c.
func (s SliceInt64) Make(l, c int) Slice {
	return make(SliceInt64, l, c)
}

// Set sets the s element at index i to v.
func (s SliceInt64) Set(i int, v interface{}) {
	s[i] = v.(int64)
}

// Slice returns the slice of s from indexes i to j.
func (s SliceInt64) Slice(i, j int) Slice {
	return s[i:j]
}

// SliceCap returns the slice of s from indexes i to j with capacity c.
func (s SliceInt64) SliceCap(i, j, c int) Slice {
	return s[i:j:c]
}

// Swap implements sort.Interface.
func (s SliceInt64) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
