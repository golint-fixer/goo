package goo

// Equatable is equatable to others of the same type.
type Equatable interface {
	// Equals returns whether e is equal.
	Equals(e Equatable) bool

	// NotEquals returns whether e is not equal.
	NotEquals(e Equatable) bool
}
