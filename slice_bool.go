package goo

// SliceBool is a slice of bool.
type SliceBool []bool

// Append appends v to s and returns the result.
func (s SliceBool) Append(v ...interface{}) Slice {
	for _, v := range v {
		s = append(s, v.(bool))
	}

	return s
}

// AppendSlice appends t to s and returns the result.
func (s SliceBool) AppendSlice(t Slice) Slice {
	return append(s, t.(SliceBool)...)
}

// Cap returns the s capacity.
func (s SliceBool) Cap() int {
	return cap(s)
}

// Copy copies t to s.
func (s SliceBool) Copy(t Slice) int {
	return copy(s, t.(SliceBool))
}

// Equals returns whether s equals v.
func (s SliceBool) Equals(v interface{}) bool {
	var t = v.(SliceBool)
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
func (s SliceBool) Get(i int) interface{} {
	return s[i]
}

// Len returns the s length.
func (s SliceBool) Len() int {
	return len(s)
}

// Make returns a new SliceBool with length l and capacity c.
func (s SliceBool) Make(l, c int) Slice {
	return make(SliceBool, l, c)
}

// Set sets the s element at index i to v.
func (s SliceBool) Set(i int, v interface{}) {
	s[i] = v.(bool)
}

// Slice returns the slice of s from indexes i to j.
func (s SliceBool) Slice(i, j int) Slice {
	return s[i:j]
}

// SliceCap returns the slice of s from indexes i to j with capacity c.
func (s SliceBool) SliceCap(i, j, c int) Slice {
	return s[i:j:c]
}

// Zero returns the zero false value of the s element type.
func (s SliceBool) Zero() interface{} {
	return false
}
