package goo

// BoolZero is the Bool zero value.
var BoolZero = Bool(false)

var _ Equatable = BoolZero

var _ Pointer = &BoolZero

var _ Value = BoolZero

// Bool is a bool.
type Bool bool

// And returns the conjunction of b and other.
func (b Bool) And(other Bool) Bool {
	return b && other
}

// Dereference implements Pointer.
func (b *Bool) Dereference() Value {
	return *b
}

// Equals implements Equatable.
func (b Bool) Equals(other Equatable) bool {
	return b == other.(Bool)
}

// Not returns the negation of b.
func (b Bool) Not() Bool {
	return !b
}

// NotEquals implements Equatable.
func (b Bool) NotEquals(other Equatable) bool {
	return b != other.(Bool)
}

// Or returns the disjunction of b and other.
func (b Bool) Or(other Bool) Bool {
	return b || other
}

// Reference implements Value.
func (b Bool) Reference() Pointer {
	return &b
}
