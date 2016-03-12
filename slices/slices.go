package slices

import "github.com/willfaught/lang"

type SliceInterface []interface{}

var _ lang.Equatable = SliceInterface(nil)

func (s SliceInterface) Append(v ...interface{}) Slice {
	return append(s, v...)
}

func (s SliceInterface) AppendSlice(t Slice) Slice {
	return append(s, t.(SliceInterface)...)
}

func (s SliceInterface) Copy(t Slice) int {
	return copy(s, t.(SliceInterface))
}

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
		if !lang.Equal(t[i], s[i]) {
			return false
		}
	}

	return true
}

func (s SliceInterface) Get(i int) interface{} {
	return s[i]
}

func (s SliceInterface) Len() int {
	return len(s)
}

func (s SliceInterface) Make(l, c int) Slice {
	return make(SliceInterface, l, c)
}

func (s SliceInterface) Set(i int, v interface{}) {
	s[i] = v
}

func (s SliceInterface) Slice(i, j int) Slice {
	return s[i:j]
}

func (s SliceInterface) SliceCap(i, j, m int) Slice {
	return s[i:j:m]
}

func (s SliceInterface) Zero() interface{} {
	return nil
}
