package goo

// BoolZero is the Bool zero value.
var BoolZero = Bool(false)

var _ Equatable = BoolZero

// Bool is a bool.
type Bool bool

// And returns the conjunction of b and other.
func (b Bool) And(other Bool) Bool {
	return b && other
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
