package goo

// Equal returns whether a and b are equal, using Equatable if possible.
func Equal(a, b interface{}) bool {
	if a, ok := a.(Equatable); ok {
		if b, ok := b.(Equatable); ok {
			return a.Equals(b)
		}
	}

	return a == b
}

// NotEqual returns whether a and b are not equal, using Equatable if possible.
func NotEqual(a, b interface{}) bool {
	if a, ok := a.(Equatable); ok {
		if b, ok := b.(Equatable); ok {
			return a.NotEquals(b)
		}
	}

	return a != b
}

// Equatable is equatable to others of the same type.
type Equatable interface {
	// Equals returns whether e is equal.
	Equals(e Equatable) bool

	// NotEquals returns whether e is not equal.
	NotEquals(e Equatable) bool
}
