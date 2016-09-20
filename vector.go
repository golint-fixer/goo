package goo

type Vector interface {
	Append(v interface{})

	Get(i int) interface{}

	//Insert(i int, v interface{})

	Len() int

	//Remove(i int)

	Set(i int, v interface{})
}

var _ Vector = vector{}

type vector struct {
	s Slice
}

/*
func NewVector() Vector {
	return vector{s: SliceInterface{}}
}
*/

func NewVectorFor(s Slice) Vector {
	return vector{s: s}
}

func (v vector) Append(w interface{}) {
	v.s = v.s.Append(w)
}

func (v vector) Get(i int) interface{} {
	return v.s.Get(i)
}

/*
func (v vector) Insert(i int, w interface{}) {
	v.s = v.s.Append(v.s.Zero())

	var l = v.s.Len()

	v.s.Slice(i+1, l).Copy(v.s.Slice(i, l))
	v.s.Set(i, w)
}
*/

func (v vector) Len() int {
	return v.s.Len()
}

/*
func (v vector) Remove(i int) {
	var l = v.s.Len()

	v.s = v.s.Slice(0, i).AppendSlice(v.s.Slice(i+1, l)).Append(v.s.Zero())
}
*/

func (v vector) Set(i int, w interface{}) {
	v.s.Set(i, w)
}
