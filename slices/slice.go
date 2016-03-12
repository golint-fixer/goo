package slices

import "github.com/willfaught/lang"

var _ lang.Iterator = &iterator{}

type Slice interface {
	Append(v ...interface{}) Slice

	AppendSlice(s Slice) Slice

	Cap() int

	Copy(s Slice) int

	Get(i int) interface{}

	Len() int

	Make(l, c int) Slice

	Set(i int, v interface{})

	Slice(i, j int) Slice

	SliceCap(i, j, m int) Slice

	Zero() interface{}
}

var _ Slice = SliceInterface(nil)

func Iterator(s Slice) lang.Iterator {
	return &iterator{n: s.Len(), s: s}
}

type iterator struct {
	i int

	n int

	s Slice
}

func (i iterator) More() bool {
	return i.i < i.n
}

func (i *iterator) Next() interface{} {
	var v = i.s.Get(i.i)

	i.i++

	return v
}
