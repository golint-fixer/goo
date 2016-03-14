package stack

import "github.com/willfaught/lang/slices"

type Stack interface {
	Len() int

	Peek() interface{}

	Pop() interface{}

	Push(v interface{})
}

var _ Stack = stack{}

type stack struct {
	s slices.Slice
}

func NewStack() Stack {
	return stack{s: slices.SliceInterface{}}
}

func NewStackFor(s slices.Slice) Stack {
	return stack{s: s}
}

func (s stack) Len() int {
	return s.s.Len()
}

func (s stack) Peek() interface{} {
	return s.s.Get(s.s.Len() - 1)
}

func (s stack) Pop() interface{} {
	var v interface{}

	s.s, v = slices.Pop(s.s)

	return v
}

func (s stack) Push(v interface{}) {
	s.s = slices.Push(s.s, v)
}
