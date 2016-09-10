package queue

import "github.com/willfaught/goo/data/slice"

type Queue interface {
	Add(v interface{})

	Len() int

	Peek() interface{}

	Remove() interface{}
}

var _ Queue = queue{}

type queue struct {
	s slice.Slice
}

func NewQueue() Queue {
	return queue{s: slice.SliceInterface{}}
}

func NewQueueFor(s slice.Slice) Queue {
	return queue{s: s}
}

func (q queue) Add(v interface{}) {
	q.s = slice.Enqueue(q.s, v)
}

func (q queue) Len() int {
	return q.s.Len()
}

func (q queue) Peek() interface{} {
	return q.s.Get(0)
}

func (q queue) Remove() interface{} {
	var v interface{}

	q.s, v = slice.Dequeue(q.s)

	return v
}
