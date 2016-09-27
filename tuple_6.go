package goo

var _ Equatable = Tuple6{}

type Tuple6 [6]interface{}

func (t Tuple6) Equals(e Equatable) bool {
	return t == e.(Tuple6)
}

func (t Tuple6) NotEquals(e Equatable) bool {
	return t != e.(Tuple6)
}

func (t Tuple6) First() interface{} {
	return t[0]
}

func (t Tuple6) Second() interface{} {
	return t[1]
}

func (t Tuple6) Third() interface{} {
	return t[2]
}

func (t Tuple6) Fourth() interface{} {
	return t[3]
}

func (t Tuple6) Fifth() interface{} {
	return t[4]
}

func (t Tuple6) Sixth() interface{} {
	return t[5]
}
