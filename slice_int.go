package goo

import "sort"

// SliceInt is a slice of int.
type SliceInt []int

var _ Slice = SliceInt(nil)

var _ sort.Interface = SliceInt(nil)

// Append appends v to s and returns the result.
func (s SliceInt) Append(v ...interface{}) Slice {
	for _, v := range v {
		s = append(s, v.(int))
	}

	return s
}

// AppendSlice appends t to s and returns the result.
func (s SliceInt) AppendSlice(t Slice) Slice {
	return append(s, t.(SliceInt)...)
}

// Cap returns the s capacity.
func (s SliceInt) Cap() int {
	return cap(s)
}

// Copy copies t to s.
func (s SliceInt) Copy(t Slice) int {
	return copy(s, t.(SliceInt))
}

// Equals returns whether s equals v.
func (s SliceInt) Equals(v interface{}) bool {
	var t = v.(SliceInt)
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
func (s SliceInt) Get(i int) interface{} {
	return s[i]
}

// Len returns the s length.
func (s SliceInt) Len() int {
	return len(s)
}

// Less implements sort.Interface.
func (s SliceInt) Less(i, j int) bool {
	return s[i] < s[j]
}

// Make returns a new SliceInt with length l and capacity c.
func (s SliceInt) Make(l, c int) Slice {
	return make(SliceInt, l, c)
}

// Set sets the s element at index i to v.
func (s SliceInt) Set(i int, v interface{}) {
	s[i] = v.(int)
}

// Slice returns the slice of s from indexes i to j.
func (s SliceInt) Slice(i, j int) Slice {
	return s[i:j]
}

// SliceCap returns the slice of s from indexes i to j with capacity c.
func (s SliceInt) SliceCap(i, j, c int) Slice {
	return s[i:j:c]
}

// Swap implements sort.Interface.
func (s SliceInt) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
