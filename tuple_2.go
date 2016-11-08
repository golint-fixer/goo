package goo

var _ Equatable = Tuple2{}

// Tuple2 is a tuple of 2 elements.
type Tuple2 [2]interface{}

// Equals implements Equatable.
func (t Tuple2) Equals(e Equatable) bool {
	return t == e.(Tuple2)
}

// NotEquals implements Equatable.
func (t Tuple2) NotEquals(e Equatable) bool {
	return t != e.(Tuple2)
}

// First returns the element at index 0.
func (t Tuple2) First() interface{} {
	return t[0]
}

// Second returns the element at index 1.
func (t Tuple2) Second() interface{} {
	return t[1]
}
