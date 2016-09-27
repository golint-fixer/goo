package goo

var _ Equatable = Tuple4{}

type Tuple4 [4]interface{}

func (t Tuple4) Equals(e Equatable) bool {
	return t == e.(Tuple4)
}

func (t Tuple4) NotEquals(e Equatable) bool {
	return t != e.(Tuple4)
}

func (t Tuple4) First() interface{} {
	return t[0]
}

func (t Tuple4) Second() interface{} {
	return t[1]
}

func (t Tuple4) Third() interface{} {
	return t[2]
}

func (t Tuple4) Fourth() interface{} {
	return t[3]
}
