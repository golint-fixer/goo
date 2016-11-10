package goo

// Tuple7 is a tuple of 7 elements.
type Tuple7 [7]interface{}

// Dereference implements Equatable.
func (t *Tuple7) Dereference() Value {
	return *t
}

// Equals implements Equatable.
func (t Tuple7) Equals(other Equatable) bool {
	return t == other.(Tuple7)
}

// NotEquals implements Equatable.
func (t Tuple7) NotEquals(other Equatable) bool {
	return t != other.(Tuple7)
}

// Reference implements Equatable.
func (t Tuple7) Reference() Pointer {
	return &t
}

// First returns the element at index 0.
func (t Tuple7) First() interface{} {
	return t[0]
}

// Second returns the element at index 1.
func (t Tuple7) Second() interface{} {
	return t[1]
}

// Third returns the element at index 2.
func (t Tuple7) Third() interface{} {
	return t[2]
}

// Fourth returns the element at index 3.
func (t Tuple7) Fourth() interface{} {
	return t[3]
}

// Fifth returns the element at index 4.
func (t Tuple7) Fifth() interface{} {
	return t[4]
}

// Sixth returns the element at index 5.
func (t Tuple7) Sixth() interface{} {
	return t[5]
}

// Seventh returns the element at index 6.
func (t Tuple7) Seventh() interface{} {
	return t[6]
}
