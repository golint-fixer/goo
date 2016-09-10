package heap

import (
	"sort"

	"github.com/willfaught/goo/data/slice"
)

type Heap interface {
	Add(v interface{})

	Len() int

	Peek() interface{}

	Remove() interface{}
}

var _ Heap = heap{}

type heap struct {
	i sort.Interface

	s slice.Slice
}

func NewHeap(s slice.Slice) Heap {
	return heap{i: s.(sort.Interface), s: s}
}

func (h heap) Add(v interface{}) {
	h.s = h.s.Append(v)

	var parent int

	for child := h.s.Len() - 1; child > 0; child = parent {
		parent = (child - 1) / 2

		if h.i.Less(parent, child) {
			break
		}

		h.i.Swap(parent, child)
	}
}

func (h heap) Len() int {
	return h.s.Len()
}

func (h heap) Peek() interface{} {
	return h.s.Get(0)
}

func (h heap) Remove() interface{} {
	var v, l = h.s.Get(0), h.s.Len()

	if l == 1 {
		h.s = h.s.Slice(0, 0)

		return v
	}

	h.s.Set(0, h.s.Get(h.s.Len()-1))
	h.s.Set(h.s.Len()-1, h.s.Zero())
	h.s = h.s.Slice(0, h.s.Len()-1)

	for parent := 0; ; {
		var left = parent*2 + 1

		if left >= l {
			break
		}

		var right = left + 1
		var child int

		if right >= l || h.i.Less(left, right) {
			child = left
		} else {
			child = right
		}

		if h.i.Less(parent, child) {
			break
		}

		h.i.Swap(parent, child)
		parent = child
	}

	return v
}
