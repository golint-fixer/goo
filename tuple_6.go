package goo

var _ Equatable = Tuple6{}

// Tuple6 is a tuple of 6 elements.
type Tuple6 [6]interface{}

// Equals implements Equatable.
func (t Tuple6) Equals(e Equatable) bool {
	return t == e.(Tuple6)
}

// NotEquals implements Equatable.
func (t Tuple6) NotEquals(e Equatable) bool {
	return t != e.(Tuple6)
}

// First returns the element at index 0.
func (t Tuple6) First() interface{} {
	return t[0]
}

// Second returns the element at index 1.
func (t Tuple6) Second() interface{} {
	return t[1]
}

// Third returns the element at index 2.
func (t Tuple6) Third() interface{} {
	return t[2]
}

// Fourth returns the element at index 3.
func (t Tuple6) Fourth() interface{} {
	return t[3]
}

// Fifth returns the element at index 4.
func (t Tuple6) Fifth() interface{} {
	return t[4]
}

// Sixth returns the element at index 5.
func (t Tuple6) Sixth() interface{} {
	return t[5]
}
