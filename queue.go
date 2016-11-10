package goo

type Queue interface {
	Add(v interface{})

	Len() int

	Peek() interface{}

	Remove() interface{}
}

var _ Queue = queue{}

type queue struct {
	s Slice
}

/*
func NewQueue() Queue {
	return queue{s: SliceInterface{}}
}
*/

func NewQueueFor(s Slice) Queue {
	return queue{s: s}
}

func (q queue) Add(v interface{}) {
	q.s = SliceEnqueue(q.s, v)
}

func (q queue) Len() int {
	return q.s.Len()
}

func (q queue) Peek() interface{} {
	return q.s.Get(0)
}

func (q queue) Remove() interface{} {
	var v interface{}

	q.s, v = SliceDequeue(q.s)

	return v
}
