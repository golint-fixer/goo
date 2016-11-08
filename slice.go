package goo

var _ Iterator = &sliceIterator{}

// Slice is a slice.
type Slice interface {
	Equatable

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

	// GetRange returns a Slice from index i to index j.
	GetRange(i, j int) Slice

	// GetRangeCap returns a Slice from index i to index j with capacity c.
	GetRangeCap(i, j, c int) Slice

	// Len returns the length.
	Len() int

	// Make returns a new Slice with length l and capacity c.
	Make(l, c int) Slice

	// Set sets the element at index i to v.
	Set(i int, v interface{})
}

// TODO: All, And, Any, Combine, CombineRepeats, Concat, Contains, Cut,
// Difference, Drop, DropEnd, DropWith, DropWithEnd, Each, Expand, Extend,
// Filter, FindIndex, FindIndexEnd, FindIndexes, First, Foldl, Foldl1, Foldr,
// Foldr1, Greatest, Group, Head, Index, Indexes, Infix, Init, Inits, Insert,
// Insert, Intersect, Intersperse, Last, Least, Map, Or, Partition, Permute,
// Prefix, Product, Remove, RemoveAll, RemoveFast, RemoveFirst, RemoveLast,
// RemoveRange, Replicate, Replicate, Reverse, Sort, SortStable, Split, Subset,
// Subslice, Subslices, Suffix, Sum, Tail, Tails, Take, TakeEnd, TakeWith,
// TakeWithEnd, Transpose, Union, Zip, Zip3-7, ZipWith, ZipWith3-7

func Dequeue(s Slice) (Slice, interface{}) {
	return s.GetRange(1, s.Len()-1), s.Get(0)
}

func Enqueue(s Slice, v interface{}) Slice {
	return s.Append(v)
}

func SliceIterator(s Slice) Iterator {
	return &sliceIterator{n: s.Len(), s: s}
}

func Pop(s Slice) (Slice, interface{}) {
	var l = s.Len()

	return s.GetRange(0, l-2), s.Get(l - 1)
}

func Push(s Slice, v interface{}) Slice {
	return s.Append(v)
}

type sliceIterator struct {
	i int

	n int

	s Slice
}

func (i sliceIterator) More() bool {
	return i.i < i.n
}

func (i *sliceIterator) Next() interface{} {
	var v = i.s.Get(i.i)

	i.i++

	return v
}
