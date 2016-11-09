package goo

import "sort"

// SliceRuneZero is the SliceRune zero value.
var SliceRuneZero = SliceRune(nil)

var _ Slice = SliceRune(nil)

var _ sort.Interface = SliceRune(nil)

// SliceRune is a slice of rune.
type SliceRune []rune

// Append implements Slice.
func (s SliceRune) Append(v ...interface{}) Slice {
	for _, v := range v {
		s = append(s, v.(rune))
	}

	return s
}

// AppendSlice implements Slice.
func (s SliceRune) AppendSlice(other Slice) Slice {
	return append(s, other.(SliceRune)...)
}

// Cap implements Slice.
func (s SliceRune) Cap() int {
	return cap(s)
}

// Copy implements Slice.
func (s SliceRune) Copy(other Slice) int {
	return copy(s, other.(SliceRune))
}

// Equals implements Slice.
func (s SliceRune) Equals(other Equatable) bool {
	var t = other.(SliceRune)

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
func (s SliceRune) Get(i int) interface{} {
	return s[i]
}

// GetRange implements Slice.
func (s SliceRune) GetRange(i, j int) Slice {
	return s[i:j]
}

// GetRangeCap implements Slice.
func (s SliceRune) GetRangeCap(i, j, c int) Slice {
	return s[i:j:c]
}

// Len implements Slice.
func (s SliceRune) Len() int {
	return len(s)
}

// Less implements sort.Interface.
func (s SliceRune) Less(i, j int) bool {
	return s[i] < s[j]
}

// Make implements Slice.
func (s SliceRune) Make(l, c int) Slice {
	return make(SliceRune, l, c)
}

// NotEquals implements Slice.
func (s SliceRune) NotEquals(other Equatable) bool {
	return !s.Equals(other)
}

// Set implements Slice.
func (s SliceRune) Set(i int, v interface{}) {
	s[i] = v.(rune)
}

// Swap implements sort.Interface.
func (s SliceRune) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
