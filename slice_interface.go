package goo

// SliceInterfaceZero is the SliceInterface zero value.
var SliceInterfaceZero = SliceInterface(nil)

var _ Slice = SliceInterface(nil)

// SliceInterface is a slice of interface{}.
type SliceInterface []interface{}

// Append implements Slice.
func (s SliceInterface) Append(v ...interface{}) Slice {
	for _, v := range v {
		s = append(s, v.(interface{}))
	}

	return s
}

// AppendSlice implements Slice.
func (s SliceInterface) AppendSlice(other Slice) Slice {
	return append(s, other.(SliceInterface)...)
}

// Cap implements Slice.
func (s SliceInterface) Cap() int {
	return cap(s)
}

// Copy implements Slice.
func (s SliceInterface) Copy(other Slice) int {
	return copy(s, other.(SliceInterface))
}

// Dereference implements Slice.
func (s *SliceInterface) Dereference() Value {
	return *s
}

// Equals implements Slice.
func (s SliceInterface) Equals(other Equatable) bool {
	var t = other.(SliceInterface)

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
func (s SliceInterface) Get(i int) interface{} {
	return s[i]
}

// GetRange implements Slice.
func (s SliceInterface) GetRange(i, j int) Slice {
	return s[i:j]
}

// GetRangeCap implements Slice.
func (s SliceInterface) GetRangeCap(i, j, c int) Slice {
	return s[i:j:c]
}

// Len implements Slice.
func (s SliceInterface) Len() int {
	return len(s)
}

// Make implements Slice.
func (s SliceInterface) Make(l, c int) Slice {
	return make(SliceInterface, l, c)
}

// NotEquals implements Slice.
func (s SliceInterface) NotEquals(other Equatable) bool {
	return !s.Equals(other)
}

// Reference implements Slice.
func (s SliceInterface) Reference() Pointer {
	return &s
}

// Set implements Slice.
func (s SliceInterface) Set(i int, v interface{}) {
	s[i] = v.(interface{})
}
