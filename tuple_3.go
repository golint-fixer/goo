package goo

var _ Equatable = Tuple3{}

// Tuple3 is a tuple of 3 elements.
type Tuple3 [3]interface{}

// Equals implements Equatable.
func (t Tuple3) Equals(e Equatable) bool {
	return t == e.(Tuple3)
}

// NotEquals implements Equatable.
func (t Tuple3) NotEquals(e Equatable) bool {
	return t != e.(Tuple3)
}

// First returns the element at index 0.
func (t Tuple3) First() interface{} {
	return t[0]
}

// Second returns the element at index 1.
func (t Tuple3) Second() interface{} {
	return t[1]
}

// Third returns the element at index 2.
func (t Tuple3) Third() interface{} {
	return t[2]
}
