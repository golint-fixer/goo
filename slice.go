package goo

var _ Iterator = &sliceIterator{}

// Slice is a slice.
type Slice interface {
	Equatable
	Value

	// Append appends v and returns the result.
	Append(v ...interface{}) Slice

	// Append appends s and returns the result.
	AppendSlice(s Slice) Slice

	// Cap returns the capacity.
	Cap() int

	// Copy copies from s and returns the number of elements copied.
	Copy(s Slice) int

	// Get returns the element at index i.
	Get(i int) interface{}

	// GetRange returns a slice from index i to index j.
	GetRange(i, j int) Slice

	// GetRangeCap returns a slice from index i to index j with capacity c.
	GetRangeCap(i, j, c int) Slice

	// Len returns the length.
	Len() int

	// Make returns a new slice with length l and capacity c.
	Make(l, c int) Slice

	// Set sets the element at index i to v.
	Set(i int, v interface{})
}
