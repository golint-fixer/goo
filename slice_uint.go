package goo

import "sort"

// SliceUint is a slice of uint.
type SliceUint []uint

var _ Slice = SliceUint(nil)

var _ sort.Interface = SliceUint(nil)

// Append appends v to s and returns the result.
func (s SliceUint) Append(v ...interface{}) Slice {
	for _, v := range v {
		s = append(s, v.(uint))
	}

	return s
}

// AppendSlice appends t to s and returns the result.
func (s SliceUint) AppendSlice(t Slice) Slice {
	return append(s, t.(SliceUint)...)
}

// Cap returns the s capacity.
func (s SliceUint) Cap() int {
	return cap(s)
}

// Copy copies t to s.
func (s SliceUint) Copy(t Slice) int {
	return copy(s, t.(SliceUint))
}

// Equals returns whether s equals v.
func (s SliceUint) Equals(v interface{}) bool {
	var t = v.(SliceUint)
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
func (s SliceUint) Get(i int) interface{} {
	return s[i]
}

// GetRange returns the slice of s from indexes i to j.
func (s SliceUint) GetRange(i, j int) Slice {
	return s[i:j]
}

// GetRangeCap returns the slice of s from indexes i to j with capacity c.
func (s SliceUint) GetRangeCap(i, j, c int) Slice {
	return s[i:j:c]
}

// Len returns the s length.
func (s SliceUint) Len() int {
	return len(s)
}

// Less implements sort.Interface.
func (s SliceUint) Less(i, j int) bool {
	return s[i] < s[j]
}

// Make returns a new SliceUint with length l and capacity c.
func (s SliceUint) Make(l, c int) Slice {
	return make(SliceUint, l, c)
}

// Set sets the s element at index i to v.
func (s SliceUint) Set(i int, v interface{}) {
	s[i] = v.(uint)
}

// Swap implements sort.Interface.
func (s SliceUint) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
