package goo

type Comparable interface {
	Equatable

	Greater(c Comparable) bool

	GreaterEqual(c Comparable) bool

	Less(c Comparable) bool

	LessEqual(c Comparable) bool
}

func Equal(a, b interface{}) bool {
	if a, ok := a.(Equatable); ok {
		if b, ok := a.(Equatable); ok {
			return a.Equals(b)
		}
	}

	return false
}