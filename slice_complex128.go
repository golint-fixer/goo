package goo

// SliceComplex128 is a slice of complex128.
type SliceComplex128 []complex128

var _ Slice = SliceComplex128(nil)

// Append appends v to s and returns the result.
func (s SliceComplex128) Append(v ...interface{}) Slice {
	for _, v := range v {
		s = append(s, v.(complex128))
	}

	return s
}

// AppendSlice appends t to s and returns the result.
func (s SliceComplex128) AppendSlice(t Slice) Slice {
	return append(s, t.(SliceComplex128)...)
}

// Cap returns the s capacity.
func (s SliceComplex128) Cap() int {
	return cap(s)
}

// Copy copies t to s.
func (s SliceComplex128) Copy(t Slice) int {
	return copy(s, t.(SliceComplex128))
}

// Equals returns whether s equals v.
func (s SliceComplex128) Equals(v interface{}) bool {
	var t = v.(SliceComplex128)
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
func (s SliceComplex128) Get(i int) interface{} {
	return s[i]
}

// Len returns the s length.
func (s SliceComplex128) Len() int {
	return len(s)
}

// Make returns a new SliceComplex128 with length l and capacity c.
func (s SliceComplex128) Make(l, c int) Slice {
	return make(SliceComplex128, l, c)
}

// Set sets the s element at index i to v.
func (s SliceComplex128) Set(i int, v interface{}) {
	s[i] = v.(complex128)
}

// Slice returns the slice of s from indexes i to j.
func (s SliceComplex128) Slice(i, j int) Slice {
	return s[i:j]
}

// SliceCap returns the slice of s from indexes i to j with capacity c.
func (s SliceComplex128) SliceCap(i, j, c int) Slice {
	return s[i:j:c]
}
