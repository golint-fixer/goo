package goo

// Tuple2 is a tuple of 2 elements.
type Tuple2 [2]interface{}

// Dereference implements Equatable.
func (t *Tuple2) Dereference() Value {
	return *t
}

// Equals implements Equatable.
func (t Tuple2) Equals(other Equatable) bool {
	return t == other.(Tuple2)
}

// NotEquals implements Equatable.
func (t Tuple2) NotEquals(other Equatable) bool {
	return t != other.(Tuple2)
}

// Reference implements Equatable.
func (t Tuple2) Reference() Pointer {
	return &t
}

// First returns the element at index 0.
func (t Tuple2) First() interface{} {
	return t[0]
}

// Second returns the element at index 1.
func (t Tuple2) Second() interface{} {
	return t[1]
}
