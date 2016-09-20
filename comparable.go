package goo

type Comparable interface {
	Equatable

	Greater(c Comparable) bool

	GreaterEqual(c Comparable) bool

	Less(c Comparable) bool

	LessEqual(c Comparable) bool
}
