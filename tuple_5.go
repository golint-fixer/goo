package goo

var _ Equatable = Tuple5{}

// Tuple5 is a tuple of 5 elements.
type Tuple5 [5]interface{}

// Equals implements Equatable.
func (t Tuple5) Equals(e Equatable) bool {
	return t == e.(Tuple5)
}

// NotEquals implements Equatable.
func (t Tuple5) NotEquals(e Equatable) bool {
	return t != e.(Tuple5)
}

// First returns the element at index 0.
func (t Tuple5) First() interface{} {
	return t[0]
}

// Second returns the element at index 1.
func (t Tuple5) Second() interface{} {
	return t[1]
}

// Third returns the element at index 2.
func (t Tuple5) Third() interface{} {
	return t[2]
}

// Fourth returns the element at index 3.
func (t Tuple5) Fourth() interface{} {
	return t[3]
}

// Fifth returns the element at index 4.
func (t Tuple5) Fifth() interface{} {
	return t[4]
}
