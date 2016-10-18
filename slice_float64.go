package goo

import "sort"

// SliceFloat64 is a slice of float64.
type SliceFloat64 []float64

var _ Slice = SliceFloat64(nil)

var _ sort.Interface = SliceFloat64(nil)

// Append appends v to s and returns the result.
func (s SliceFloat64) Append(v ...interface{}) Slice {
	for _, v := range v {
		s = append(s, v.(float64))
	}

	return s
}

// AppendSlice appends t to s and returns the result.
func (s SliceFloat64) AppendSlice(t Slice) Slice {
	return append(s, t.(SliceFloat64)...)
}

// Cap returns the s capacity.
func (s SliceFloat64) Cap() int {
	return cap(s)
}

// Copy copies t to s.
func (s SliceFloat64) Copy(t Slice) int {
	return copy(s, t.(SliceFloat64))
}

// Equals returns whether s equals v.
func (s SliceFloat64) Equals(v interface{}) bool {
	var t = v.(SliceFloat64)
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
func (s SliceFloat64) Get(i int) interface{} {
	return s[i]
}

// GetRange returns the slice of s from indexes i to j.
func (s SliceFloat64) GetRange(i, j int) Slice {
	return s[i:j]
}

// GetRangeCap returns the slice of s from indexes i to j with capacity c.
func (s SliceFloat64) GetRangeCap(i, j, c int) Slice {
	return s[i:j:c]
}

// Len returns the s length.
func (s SliceFloat64) Len() int {
	return len(s)
}

// Less implements sort.Interface.
func (s SliceFloat64) Less(i, j int) bool {
	return s[i] < s[j]
}

// Make returns a new SliceFloat64 with length l and capacity c.
func (s SliceFloat64) Make(l, c int) Slice {
	return make(SliceFloat64, l, c)
}

// Set sets the s element at index i to v.
func (s SliceFloat64) Set(i int, v interface{}) {
	s[i] = v.(float64)
}

// Swap implements sort.Interface.
func (s SliceFloat64) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
