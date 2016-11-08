package goo

// Comparable is comparable to others of the same type.
type Comparable interface {
	Equatable

	Greater(c Comparable) bool

	GreaterEqual(c Comparable) bool

	Less(c Comparable) bool

	LessEqual(c Comparable) bool
}
