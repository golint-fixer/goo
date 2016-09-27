package goo

var _ Equatable = Tuple5{}

type Tuple5 [5]interface{}

func (t Tuple5) Equals(e Equatable) bool {
	return t == e.(Tuple5)
}

func (t Tuple5) NotEquals(e Equatable) bool {
	return t != e.(Tuple5)
}

func (t Tuple5) First() interface{} {
	return t[0]
}

func (t Tuple5) Second() interface{} {
	return t[1]
}

func (t Tuple5) Third() interface{} {
	return t[2]
}

func (t Tuple5) Fourth() interface{} {
	return t[3]
}

func (t Tuple5) Fifth() interface{} {
	return t[4]
}
