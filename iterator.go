package lang

type Iterator interface {
	More() bool

	Next() interface{}
}

func Repeat(v interface{}) Iterator {
	return repeater{v: v}
}

func Replicate(v interface{}, n int) Iterator {
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
