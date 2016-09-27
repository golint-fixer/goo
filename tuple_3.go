package goo

var _ Equatable = Tuple3{}

type Tuple3 [3]interface{}

func (t Tuple3) Equals(e Equatable) bool {
	return t == e.(Tuple3)
}

func (t Tuple3) NotEquals(e Equatable) bool {
	return t != e.(Tuple3)
}

func (t Tuple3) First() interface{} {
	return t[0]
}

func (t Tuple3) Second() interface{} {
	return t[1]
}

func (t Tuple3) Third() interface{} {
	return t[2]
}
