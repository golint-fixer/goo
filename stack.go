package goo

type Stack interface {
	Len() int

	Peek() interface{}

	Pop() interface{}

	Push(v interface{})
}

var _ Stack = stack{}

type stack struct {
	s Slice
}

/*
func NewStack() Stack {
	return stack{s: SliceInterface{}}
}
*/

func NewStackFor(s Slice) Stack {
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

	s.s, v = SlicePop(s.s)

	return v
}

func (s stack) Push(v interface{}) {
	s.s = SlicePush(s.s, v)
}
