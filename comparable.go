package goo

// Comparable is comparable to others of the same type.
type Comparable interface {
	Equatable

	// Greater returns whether c is less or equal.
	Greater(c Comparable) bool

	// GreaterEqual returns whether c is less.
	GreaterEqual(c Comparable) bool

	// Less returns whether c is greater or equal.
	Less(c Comparable) bool

	// LessEqual returns whether c is greater.
	LessEqual(c Comparable) bool
}
