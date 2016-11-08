package goo

// Iterator iterates a sequence of elements.
type Iterator interface {
	// More returns whether there is a next element.
	More() bool

	// Next returns the next element.
	Next() interface{}
}

// RepeatIterator returns an Iterator that repeats v forever.
func RepeatIterator(v interface{}) Iterator {
	return repeater{v: v}
}

// ReplicateIterator returns an Iterator that contains n v.
func ReplicateIterator(v interface{}, n int) Iterator {
	return &replicator{n: n, v: v}
}

type repeater struct {
	v interface{}
}

func (r repeater) More() bool {
	return true
}

func (r repeater) Next() interface{} {
	return r.v
}

type replicator struct {
	i int
	n int
	v interface{}
}

func (r replicator) More() bool {
	return r.i < r.n
}

func (r *replicator) Next() interface{} {
	r.i++

	return r.v
}
