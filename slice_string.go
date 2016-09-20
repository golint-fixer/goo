package goo

import "sort"

// SliceString is a slice of string.
type SliceString []string

var _ Slice = SliceString(nil)

var _ sort.Interface = SliceString(nil)

// Append appends v to s and returns the result.
func (s SliceString) Append(v ...interface{}) Slice {
	for _, v := range v {
		s = append(s, v.(string))
	}

	return s
}

// AppendSlice appends t to s and returns the result.
func (s SliceString) AppendSlice(t Slice) Slice {
	return append(s, t.(SliceString)...)
}

// Cap returns the s capacity.
func (s SliceString) Cap() int {
	return cap(s)
}

// Copy copies t to s.
func (s SliceString) Copy(t Slice) int {
	return copy(s, t.(SliceString))
}

// Equals returns whether s equals v.
func (s SliceString) Equals(v interface{}) bool {
	var t = v.(SliceString)
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
func (s SliceString) Get(i int) interface{} {
	return s[i]
}

// Len returns the s length.
func (s SliceString) Len() int {
	return len(s)
}

// Less implements sort.Interface.
func (s SliceString) Less(i, j int) bool {
	return s[i] < s[j]
}

// Make returns a new SliceString with length l and capacity c.
func (s SliceString) Make(l, c int) Slice {
	return make(SliceString, l, c)
}

// Set sets the s element at index i to v.
func (s SliceString) Set(i int, v interface{}) {
	s[i] = v.(string)
}

// Slice returns the slice of s from indexes i to j.
func (s SliceString) Slice(i, j int) Slice {
	return s[i:j]
}

// SliceCap returns the slice of s from indexes i to j with capacity c.
func (s SliceString) SliceCap(i, j, c int) Slice {
	return s[i:j:c]
}

// Swap implements sort.Interface.
func (s SliceString) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
