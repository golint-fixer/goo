package slices

type Slice interface {
	Append(v ...interface{}) Slice

	AppendSlice(s Slice) Slice

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
