package queue

import "github.com/willfaught/lang/slices"

type Queue interface {
	Add(v interface{})

	Len() int

	Peek() interface{}

	Remove() interface{}
}

var _ Queue = queue{}

type queue struct {
	s slices.Slice
}

func NewQueue() Queue {
	return queue{s: slices.SliceInterface{}}
}

func NewQueueFor(s slices.Slice) Queue {
	return queue{s: s}
}

func (q queue) Add(v interface{}) {
	q.s = slices.Enqueue(q.s, v)
}

func (q queue) Len() int {
	return q.s.Len()
}

func (q queue) Peek() interface{} {
	return q.s.Get(0)
}

func (q queue) Remove() interface{} {
	var v interface{}

	q.s, v = slices.Dequeue(q.s)

	return v
}
