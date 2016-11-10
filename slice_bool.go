package goo

// SliceBoolZero is the SliceBool zero value.
var SliceBoolZero = SliceBool(nil)

var _ Pointer = &SliceBool{}

var _ Slice = SliceBool(nil)

// SliceBool is a slice of bool.
type SliceBool []bool

// Append implements Slice.
func (s SliceBool) Append(v ...interface{}) Slice {
	for _, v := range v {
		s = append(s, v.(bool))
	}

	return s
}

// AppendSlice implements Slice.
func (s SliceBool) AppendSlice(other Slice) Slice {
	return append(s, other.(SliceBool)...)
}

// Cap implements Slice.
func (s SliceBool) Cap() int {
	return cap(s)
}

// Copy implements Slice.
func (s SliceBool) Copy(other Slice) int {
	return copy(s, other.(SliceBool))
}

// Dereference implements Slice.
func (s *SliceBool) Dereference() Value {
	return *s
}

// Equals implements Slice.
func (s SliceBool) Equals(other Equatable) bool {
	var t = other.(SliceBool)

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
func (s SliceBool) Get(i int) interface{} {
	return s[i]
}

// GetRange implements Slice.
func (s SliceBool) GetRange(i, j int) Slice {
	return s[i:j]
}

// GetRangeCap implements Slice.
func (s SliceBool) GetRangeCap(i, j, c int) Slice {
	return s[i:j:c]
}

// Len implements Slice.
func (s SliceBool) Len() int {
	return len(s)
}

// Make implements Slice.
func (s SliceBool) Make(l, c int) Slice {
	return make(SliceBool, l, c)
}

// NotEquals implements Slice.
func (s SliceBool) NotEquals(other Equatable) bool {
	return !s.Equals(other)
}

// Reference implements Slice.
func (s SliceBool) Reference() Pointer {
	return &s
}

// Set implements Slice.
func (s SliceBool) Set(i int, v interface{}) {
	s[i] = v.(bool)
}
