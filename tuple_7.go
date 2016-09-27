package goo

var _ Equatable = Tuple7{}

type Tuple7 [7]interface{}

func (t Tuple7) Equals(e Equatable) bool {
	return t == e.(Tuple7)
}

func (t Tuple7) NotEquals(e Equatable) bool {
	return t != e.(Tuple7)
}

func (t Tuple7) First() interface{} {
	return t[0]
}

func (t Tuple7) Second() interface{} {
	return t[1]
}

func (t Tuple7) Third() interface{} {
	return t[2]
}

func (t Tuple7) Fourth() interface{} {
	return t[3]
}

func (t Tuple7) Fifth() interface{} {
	return t[4]
}

func (t Tuple7) Sixth() interface{} {
	return t[5]
}

func (t Tuple7) Seventh() interface{} {
	return t[6]
}
