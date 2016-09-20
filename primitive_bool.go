package goo

var _ Equatable = Bool(false)

// BoolZero is the Bool zero value.
var BoolZero = Bool(false)

// Bool is a bool.
type Bool bool

// And returns the conjunction of b and other.
func (b Bool) And(other Bool) Bool {
	return b && other
}

// Equals implements Equatable.
func (b Bool) Equals(e Equatable) bool {
	return b == e.(Bool)
}

// Not returns the negation of b.
func (b Bool) Not() Bool {
	return !b
}

// NotEquals implements Equatable.
func (b Bool) NotEquals(e Equatable) bool {
	return b != e.(Bool)
}

// Or returns the disjunction of b and other.
func (b Bool) Or(other Bool) Bool {
	return b || other
}
