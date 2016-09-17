package goo

// SliceInterface is a slice of interface{}.
type SliceInterface []interface{}

// Append appends v to s and returns the result.
func (s SliceInterface) Append(v ...interface{}) Slice {
	for _, v := range v {
		s = append(s, v.(interface{}))
	}

	return s
}

// AppendSlice appends t to s and returns the result.
func (s SliceInterface) AppendSlice(t Slice) Slice {
	return append(s, t.(SliceInterface)...)
}

// Cap returns the s capacity.
func (s SliceInterface) Cap() int {
	return cap(s)
}

// Copy copies t to s.
func (s SliceInterface) Copy(t Slice) int {
	return copy(s, t.(SliceInterface))
}

// Equals returns whether s equals v.
func (s SliceInterface) Equals(v interface{}) bool {
	var t = v.(SliceInterface)
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
func (s SliceInterface) Get(i int) interface{} {
	return s[i]
}

// Len returns the s length.
func (s SliceInterface) Len() int {
	return len(s)
}

// Make returns a new SliceInterface with length l and capacity c.
func (s SliceInterface) Make(l, c int) Slice {
	return make(SliceInterface, l, c)
}

// Set sets the s element at index i to v.
func (s SliceInterface) Set(i int, v interface{}) {
	s[i] = v.(interface{})
}

// Slice returns the slice of s from indexes i to j.
func (s SliceInterface) Slice(i, j int) Slice {
	return s[i:j]
}

// SliceCap returns the slice of s from indexes i to j with capacity c.
func (s SliceInterface) SliceCap(i, j, c int) Slice {
	return s[i:j:c]
}

// Zero returns the zero nil value of the s element type.
func (s SliceInterface) Zero() interface{} {
	return nil
}
