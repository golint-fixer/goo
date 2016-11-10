package goo

// Tuple4 is a tuple of 4 elements.
type Tuple4 [4]interface{}

// Dereference implements Equatable.
func (t *Tuple4) Dereference() Value {
	return *t
}

// Equals implements Equatable.
func (t Tuple4) Equals(other Equatable) bool {
	return t == other.(Tuple4)
}

// NotEquals implements Equatable.
func (t Tuple4) NotEquals(other Equatable) bool {
	return t != other.(Tuple4)
}

// Reference implements Equatable.
func (t Tuple4) Reference() Pointer {
	return &t
}

// First returns the element at index 0.
func (t Tuple4) First() interface{} {
	return t[0]
}

// Second returns the element at index 1.
func (t Tuple4) Second() interface{} {
	return t[1]
}

// Third returns the element at index 2.
func (t Tuple4) Third() interface{} {
	return t[2]
}

// Fourth returns the element at index 3.
func (t Tuple4) Fourth() interface{} {
	return t[3]
}
