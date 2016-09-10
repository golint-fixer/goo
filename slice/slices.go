package slice

import (
	"sort"

	"github.com/willfaught/goo"
)

var (
	_ lang.Equatable = SliceBool(nil)
	_ lang.Equatable = SliceInt(nil)
	_ lang.Equatable = SliceInterface(nil)
	_ lang.Equatable = SliceRune(nil)
	_ lang.Equatable = SliceString(nil)
)

var (
	_ lang.EquatableDeep = SliceBool(nil)
	_ lang.EquatableDeep = SliceInt(nil)
	_ lang.EquatableDeep = SliceInterface(nil)
	_ lang.EquatableDeep = SliceRune(nil)
	_ lang.EquatableDeep = SliceString(nil)
)

var (
	_ sort.Interface = SliceBool(nil)
	_ sort.Interface = SliceInt(nil)
	_ sort.Interface = SliceRune(nil)
	_ sort.Interface = SliceString(nil)
)

type SliceBool []bool

func (s SliceBool) Append(v ...interface{}) Slice {
	for _, v := range v {
		s = append(s, v.(bool))
	}

	return s
}

func (s SliceBool) AppendSlice(t Slice) Slice {
	return append(s, t.(SliceBool)...)
}

func (s SliceBool) Cap() int {
	return cap(s)
}

func (s SliceBool) Copy(t Slice) int {
	return copy(s, t.(SliceBool))
}

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

func (s SliceBool) EqualsDeep(v interface{}) bool {
	var t = v.(SliceBool)
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

func (s SliceBool) Get(i int) interface{} {
	return s[i]
}

func (s SliceBool) Len() int {
	return len(s)
}

func (s SliceBool) Less(i, j int) bool {
	return !s[i] && s[j]
}

func (s SliceBool) Make(l, c int) Slice {
	return make(SliceBool, l, c)
}

func (s SliceBool) Set(i int, v interface{}) {
	s[i] = v.(bool)
}

func (s SliceBool) Slice(i, j int) Slice {
	return s[i:j]
}

func (s SliceBool) SliceCap(i, j, m int) Slice {
	return s[i:j:m]
}

func (s SliceBool) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s SliceBool) Zero() interface{} {
	return false
}

type SliceInt []int

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
		if t[i] != s[i] {
			return false
		}
	}

	return true
}

func (s SliceInt) EqualsDeep(v interface{}) bool {
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
		if t[i] != s[i] {
			return false
		}
	}

	return true
}

func (s SliceInterface) EqualsDeep(v interface{}) bool {
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
		if t[i] != s[i] {
			return false
		}
	}

	return true
}

func (s SliceRune) EqualsDeep(v interface{}) bool {
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

type SliceString []string

func (s SliceString) Append(v ...interface{}) Slice {
	for _, v := range v {
		s = append(s, v.(string))
	}

	return s
}

func (s SliceString) AppendSlice(t Slice) Slice {
	return append(s, t.(SliceString)...)
}

func (s SliceString) Cap() int {
	return cap(s)
}

func (s SliceString) Copy(t Slice) int {
	return copy(s, t.(SliceString))
}

func (s SliceString) Equals(v interface{}) bool {
	var t = v.(SliceString)
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

func (s SliceString) EqualsDeep(v interface{}) bool {
	var t = v.(SliceString)
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

func (s SliceString) Get(i int) interface{} {
	return s[i]
}

func (s SliceString) Len() int {
	return len(s)
}

func (s SliceString) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s SliceString) Make(l, c int) Slice {
	return make(SliceString, l, c)
}

func (s SliceString) Set(i int, v interface{}) {
	s[i] = v.(string)
}

func (s SliceString) Slice(i, j int) Slice {
	return s[i:j]
}

func (s SliceString) SliceCap(i, j, m int) Slice {
	return s[i:j:m]
}

func (s SliceString) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s SliceString) Zero() interface{} {
	return ""
}
