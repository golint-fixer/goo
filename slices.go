package goo

// TODO: Combine, CombineRepeats, , , Cut,
// Difference, , , , , Expand, Extend,
// , FindIndex, FindIndexEnd, FindIndexes, First, Foldl, Foldl1, Foldr,
// Foldr1, Greatest, Group, , Index, Indexes, Infix, Init, Inits, Insert,
// Insert, Intersect, Intersperse, , Least, , Partition, Permute,
// Prefix, Product, Remove, RemoveAll, RemoveFast, RemoveFirst, RemoveLast,
// RemoveRange, Replicate, Replicate, Reverse, Sort, SortStable, Split, Subset,
// Subslice, Subslices, Suffix, Sum, , Tails, Take, TakeEnd, TakeWith,
// TakeWithEnd, Transpose, Union, Zip, Zip3-7, ZipWith, ZipWith3-7

// SliceAll returns whether all s elements are true.
func SliceAll(s Slice) bool {
	for i, n := 0, s.Len(); i < n; i++ {
		if !s.Get(i).(bool) {
			return false
		}
	}

	return true
}

// SliceAny returns whether any s elements are true.
func SliceAny(s Slice) bool {
	for i, n := 0, s.Len(); i < n; i++ {
		if s.Get(i).(bool) {
			return true
		}
	}

	return false
}

// SliceContains returns whether s contains v.
func SliceContains(s Slice, v interface{}) bool {
	for i, n := 0, s.Len(); i < n; i++ {
		if Equal(s.Get(i), v) {
			return true
		}
	}

	return false
}

// SliceDequeue returns the result of popping the first s element.
func SliceDequeue(s Slice) (Slice, interface{}) {
	return s.GetRange(1, s.Len()-1), s.Get(0)
}

// SliceDrop returns s without the first n elements.
func SliceDrop(s Slice, n int) Slice {
	var l = s.Len()

	if n > l {
		n = l
	}

	return s.GetRange(l-n, l)
}

// SliceDropEnd returns s without the last n elements.
func SliceDropEnd(s Slice, n int) Slice {
	var l = s.Len()

	if n > l {
		n = l
	}

	return s.GetRange(0, l-n)
}

// SliceDropWhile returns s without the leading elements that do not satisfy f.
func SliceDropWhile(s Slice, f func(interface{}) bool) Slice {
	for i, n := 0, s.Len(); i < n; i++ {
		if !f(s.Get(i)) {
			return s.GetRange(i+1, n)
		}
	}

	return s
}

// SliceDropWhileEnd returns s without the trailing elements that do not satisfy f.
func SliceDropWhileEnd(s Slice, f func(interface{}) bool) Slice {
	for i := s.Len() - 1; i >= 0; i-- {
		if !f(s.Get(i)) {
			return s.GetRange(0, i)
		}
	}

	return s
}

// SliceEach calls f with every s element.
func SliceEach(s Slice, f func(interface{})) {
	for i, n := 0, s.Len(); i < n; i++ {
		f(s.Get(i))
	}
}

// SliceEnqueue returns the result of appending v to s.
func SliceEnqueue(s Slice, v interface{}) Slice {
	return s.Append(v)
}

// SliceFilter filters out s elements for which f returns false.
func SliceFilter(s Slice, f func(interface{}) bool) Slice {
	var t Slice

	for i, n := 0, s.Len(); i < n; i++ {
		if v := s.Get(i); f(v) {
			t = t.Append(t, v)
		}
	}

	return t
}

// SliceFirst returns the first s element.
func SliceFirst(s Slice) interface{} {
	return s.Get(0)
}

// SliceHead returns the first s element.
func SliceHead(s Slice) interface{} {
	return s.Get(0)
}

// SliceIterator returns an s Iterator.
func SliceIterator(s Slice) Iterator {
	return &sliceIterator{n: s.Len(), s: s}
}

// SliceLast returns the last s element.
func SliceLast(s Slice) interface{} {
	return s.Get(s.Len() - 1)
}

// SliceMap returns the s elements transformed by f.
func SliceMap(s Slice, f func(interface{}) interface{}) Slice {
	var t Slice

	for i, n := 0, s.Len(); i < n; i++ {
		t = t.Append(t, f(s.Get(i)))
	}

	return t
}

// SlicePop returns the result of popping the last s element.
func SlicePop(s Slice) (Slice, interface{}) {
	var l = s.Len()

	return s.GetRange(0, l-1), s.Get(l - 1)
}

// SlicePush returns the result of appending v to s.
func SlicePush(s Slice, v interface{}) Slice {
	return s.Append(v)
}

// SliceTail returns all but the first s element.
func SliceTail(s Slice) Slice {
	return s.GetRange(1, s.Len())
}

type sliceIterator struct {
	i int
	n int
	s Slice
}

func (i *sliceIterator) More() bool {
	return i.i < i.n
}

func (i *sliceIterator) Next() interface{} {
	var v = i.s.Get(i.i)

	i.i++

	return v
}
