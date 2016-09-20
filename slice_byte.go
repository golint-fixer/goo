package goo

import "sort"

// SliceByte is a slice of byte.
type SliceByte []byte

var _ Slice = SliceByte(nil)

var _ sort.Interface = SliceByte(nil)

// Append appends v to s and returns the result.
func (s SliceByte) Append(v ...interface{}) Slice {
	for _, v := range v {
		s = append(s, v.(byte))
	}

	return s
}

// AppendSlice appends t to s and returns the result.
func (s SliceByte) AppendSlice(t Slice) Slice {
	return append(s, t.(SliceByte)...)
}

// Cap returns the s capacity.
func (s SliceByte) Cap() int {
	return cap(s)
}

// Copy copies t to s.
func (s SliceByte) Copy(t Slice) int {
	return copy(s, t.(SliceByte))
}

// Equals returns whether s equals v.
func (s SliceByte) Equals(v interface{}) bool {
	var t = v.(SliceByte)
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
func (s SliceByte) Get(i int) interface{} {
	return s[i]
}

// Len returns the s length.
func (s SliceByte) Len() int {
	return len(s)
}

// Less implements sort.Interface.
func (s SliceByte) Less(i, j int) bool {
	return s[i] < s[j]
}

// Make returns a new SliceByte with length l and capacity c.
func (s SliceByte) Make(l, c int) Slice {
	return make(SliceByte, l, c)
}

// Set sets the s element at index i to v.
func (s SliceByte) Set(i int, v interface{}) {
	s[i] = v.(byte)
}

// Slice returns the slice of s from indexes i to j.
func (s SliceByte) Slice(i, j int) Slice {
	return s[i:j]
}

// SliceCap returns the slice of s from indexes i to j with capacity c.
func (s SliceByte) SliceCap(i, j, c int) Slice {
	return s[i:j:c]
}

// Swap implements sort.Interface.
func (s SliceByte) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
