package goo

// SliceComplex128Zero is the SliceComplex128 zero value.
var SliceComplex128Zero = SliceComplex128(nil)

var _ Slice = SliceComplex128(nil)

// SliceComplex128 is a slice of complex128.
type SliceComplex128 []complex128

// Append implements Slice.
func (s SliceComplex128) Append(v ...interface{}) Slice {
	for _, v := range v {
		s = append(s, v.(complex128))
	}

	return s
}

// AppendSlice implements Slice.
func (s SliceComplex128) AppendSlice(other Slice) Slice {
	return append(s, other.(SliceComplex128)...)
}

// Cap implements Slice.
func (s SliceComplex128) Cap() int {
	return cap(s)
}

// Copy implements Slice.
func (s SliceComplex128) Copy(other Slice) int {
	return copy(s, other.(SliceComplex128))
}

// Equals implements Slice.
func (s SliceComplex128) Equals(other Equatable) bool {
	var t = other.(SliceComplex128)

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
func (s SliceComplex128) Get(i int) interface{} {
	return s[i]
}

// GetRange implements Slice.
func (s SliceComplex128) GetRange(i, j int) Slice {
	return s[i:j]
}

// GetRangeCap implements Slice.
func (s SliceComplex128) GetRangeCap(i, j, c int) Slice {
	return s[i:j:c]
}

// Len implements Slice.
func (s SliceComplex128) Len() int {
	return len(s)
}

// Make implements Slice.
func (s SliceComplex128) Make(l, c int) Slice {
	return make(SliceComplex128, l, c)
}

// NotEquals implements Slice.
func (s SliceComplex128) NotEquals(other Equatable) bool {
	return !s.Equals(other)
}

// Set implements Slice.
func (s SliceComplex128) Set(i int, v interface{}) {
	s[i] = v.(complex128)
}
