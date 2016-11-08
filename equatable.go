package goo

type Equatable interface {
	Equals(e Equatable) bool

	NotEquals(e Equatable) bool
}
