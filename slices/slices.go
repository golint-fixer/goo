package slices

import (
	"sort"

	"github.com/willfaught/lang"
)

type SliceInt []int

var (
	_ lang.Equatable = SliceInt(nil)
	_ sort.Interface = SliceInt(nil)
)

func (s SliceInt) Append(v ...interface{}) Slice {
	for _, v := range v {
		s = append(s, v.(int))
	}

	return s
}

func (s SliceInt) AppendSlice(t Slice) Slice {
	return append(s, t.(SliceInt)...)
}

func (s SliceInt) Cap() int {
	return cap(s)
}

func (s SliceInt) Copy(t Slice) int {
	return copy(s, t.(SliceInt))
}

func (s SliceInt) Equals(v interface{}) bool {
	var t = v.(SliceInt)
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

func (s SliceInt) Get(i int) interface{} {
	return s[i]
}

func (s SliceInt) Len() int {
	return len(s)
}

func (s SliceInt) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s SliceInt) Make(l, c int) Slice {
	return make(SliceInt, l, c)
}

func (s SliceInt) Set(i int, v interface{}) {
	s[i] = v.(int)
}

func (s SliceInt) Slice(i, j int) Slice {
	return s[i:j]
}

func (s SliceInt) SliceCap(i, j, m int) Slice {
	return s[i:j:m]
}

func (s SliceInt) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s SliceInt) Zero() interface{} {
	return int(0)
}

type SliceInterface []interface{}

var _ lang.Equatable = SliceInterface(nil)

func (s SliceInterface) Append(v ...interface{}) Slice {
	return append(s, v...)
}

func (s SliceInterface) AppendSlice(t Slice) Slice {
	return append(s, t.(SliceInterface)...)
}

func (s SliceInterface) Cap() int {
	return cap(s)
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

type SliceRune []rune

var (
	_ lang.Equatable = SliceRune(nil)
	_ sort.Interface = SliceRune(nil)
)

func (s SliceRune) Append(v ...interface{}) Slice {
	for _, v := range v {
		s = append(s, v.(rune))
	}

	return s
}

func (s SliceRune) AppendSlice(t Slice) Slice {
	return append(s, t.(SliceRune)...)
}

func (s SliceRune) Cap() int {
	return cap(s)
}

func (s SliceRune) Copy(t Slice) int {
	return copy(s, t.(SliceRune))
}

func (s SliceRune) Equals(v interface{}) bool {
	var t = v.(SliceRune)
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

func (s SliceRune) Get(i int) interface{} {
	return s[i]
}

func (s SliceRune) Len() int {
	return len(s)
}

func (s SliceRune) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s SliceRune) Make(l, c int) Slice {
	return make(SliceRune, l, c)
}

func (s SliceRune) Set(i int, v interface{}) {
	s[i] = v.(rune)
}

func (s SliceRune) Slice(i, j int) Slice {
	return s[i:j]
}

func (s SliceRune) SliceCap(i, j, m int) Slice {
	return s[i:j:m]
}

func (s SliceRune) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s SliceRune) Zero() interface{} {
	return rune(0)
}
