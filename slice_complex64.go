package goo

// SliceComplex64 is a slice of complex64.
type SliceComplex64 []complex64

// Append appends v to s and returns the result.
func (s SliceComplex64) Append(v ...interface{}) Slice {
	for _, v := range v {
		s = append(s, v.(complex64))
	}

	return s
}

// AppendSlice appends t to s and returns the result.
func (s SliceComplex64) AppendSlice(t Slice) Slice {
	return append(s, t.(SliceComplex64)...)
}

// Cap returns the s capacity.
func (s SliceComplex64) Cap() int {
	return cap(s)
}

// Copy copies t to s.
func (s SliceComplex64) Copy(t Slice) int {
	return copy(s, t.(SliceComplex64))
}

// Equals returns whether s equals v.
func (s SliceComplex64) Equals(v interface{}) bool {
	var t = v.(SliceComplex64)
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
func (s SliceComplex64) Get(i int) interface{} {
	return s[i]
}

// Len returns the s length.
func (s SliceComplex64) Len() int {
	return len(s)
}

// Make returns a new SliceComplex64 with length l and capacity c.
func (s SliceComplex64) Make(l, c int) Slice {
	return make(SliceComplex64, l, c)
}

// Set sets the s element at index i to v.
func (s SliceComplex64) Set(i int, v interface{}) {
	s[i] = v.(complex64)
}

// Slice returns the slice of s from indexes i to j.
func (s SliceComplex64) Slice(i, j int) Slice {
	return s[i:j]
}

// SliceCap returns the slice of s from indexes i to j with capacity c.
func (s SliceComplex64) SliceCap(i, j, c int) Slice {
	return s[i:j:c]
}

// Zero returns the zero 0 value of the s element type.
func (s SliceComplex64) Zero() interface{} {
	return 0
}
