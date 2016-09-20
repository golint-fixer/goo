package goo

func Equal(a, b interface{}) bool {
	if a, ok := a.(Equatable); ok {
		if b, ok := a.(Equatable); ok {
			return a.Equals(b)
		}
	}

	return false
}

func NotEqual(a, b interface{}) bool {
	if a, ok := a.(Equatable); ok {
		if b, ok := a.(Equatable); ok {
			return a.NotEquals(b)
		}
	}

	return true
}

type Equatable interface {
	Equals(e Equatable) bool

	NotEquals(e Equatable) bool
}
