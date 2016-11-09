package goo

// SliceComplex64Zero is the SliceComplex64 zero value.
var SliceComplex64Zero = SliceComplex64(nil)

var _ Slice = SliceComplex64(nil)

// SliceComplex64 is a slice of complex64.
type SliceComplex64 []complex64

// Append implements Slice.
func (s SliceComplex64) Append(v ...interface{}) Slice {
	for _, v := range v {
		s = append(s, v.(complex64))
	}

	return s
}

// AppendSlice implements Slice.
func (s SliceComplex64) AppendSlice(other Slice) Slice {
	return append(s, other.(SliceComplex64)...)
}

// Cap implements Slice.
func (s SliceComplex64) Cap() int {
	return cap(s)
}

// Copy implements Slice.
func (s SliceComplex64) Copy(other Slice) int {
	return copy(s, other.(SliceComplex64))
}

// Equals implements Slice.
func (s SliceComplex64) Equals(other Equatable) bool {
	var t = other.(SliceComplex64)

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
func (s SliceComplex64) Get(i int) interface{} {
	return s[i]
}

// GetRange implements Slice.
func (s SliceComplex64) GetRange(i, j int) Slice {
	return s[i:j]
}

// GetRangeCap implements Slice.
func (s SliceComplex64) GetRangeCap(i, j, c int) Slice {
	return s[i:j:c]
}

// Len implements Slice.
func (s SliceComplex64) Len() int {
	return len(s)
}

// Make implements Slice.
func (s SliceComplex64) Make(l, c int) Slice {
	return make(SliceComplex64, l, c)
}

// NotEquals implements Slice.
func (s SliceComplex64) NotEquals(other Equatable) bool {
	return !s.Equals(other)
}

// Set implements Slice.
func (s SliceComplex64) Set(i int, v interface{}) {
	s[i] = v.(complex64)
}
