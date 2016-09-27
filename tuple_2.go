package goo

var _ Equatable = Tuple2{}

type Tuple2 [2]interface{}

func (t Tuple2) Equals(e Equatable) bool {
	return t == e.(Tuple2)
}

func (t Tuple2) NotEquals(e Equatable) bool {
	return t != e.(Tuple2)
}

func (t Tuple2) First() interface{} {
	return t[0]
}

func (t Tuple2) Second() interface{} {
	return t[1]
}
