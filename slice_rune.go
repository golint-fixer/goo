package goo

import "sort"

// SliceRune is a slice of rune.
type SliceRune []rune

var _ sort.Interface = SliceRune(nil)

// Append appends v to s and returns the result.
func (s SliceRune) Append(v ...interface{}) Slice {
	for _, v := range v {
		s = append(s, v.(rune))
	}

	return s
}

// AppendSlice appends t to s and returns the result.
func (s SliceRune) AppendSlice(t Slice) Slice {
	return append(s, t.(SliceRune)...)
}

// Cap returns the s capacity.
func (s SliceRune) Cap() int {
	return cap(s)
}

// Copy copies t to s.
func (s SliceRune) Copy(t Slice) int {
	return copy(s, t.(SliceRune))
}

// Equals returns whether s equals v.
func (s SliceRune) Equals(v interface{}) bool {
	var t = v.(SliceRune)
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
func (s SliceRune) Get(i int) interface{} {
	return s[i]
}

// Len returns the s length.
func (s SliceRune) Len() int {
	return len(s)
}

// Less implements sort.Interface.
func (s SliceRune) Less(i, j int) bool {
	return s[i] < s[j]
}

// Make returns a new SliceRune with length l and capacity c.
func (s SliceRune) Make(l, c int) Slice {
	return make(SliceRune, l, c)
}

// Set sets the s element at index i to v.
func (s SliceRune) Set(i int, v interface{}) {
	s[i] = v.(rune)
}

// Slice returns the slice of s from indexes i to j.
func (s SliceRune) Slice(i, j int) Slice {
	return s[i:j]
}

// SliceCap returns the slice of s from indexes i to j with capacity c.
func (s SliceRune) SliceCap(i, j, c int) Slice {
	return s[i:j:c]
}

// Swap implements sort.Interface.
func (s SliceRune) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// Zero returns the zero 0 value of the s element type.
func (s SliceRune) Zero() interface{} {
	return 0
}
