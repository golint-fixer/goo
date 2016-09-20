package goo

var _ Comparable = String("")

// StringZero is the String zero value.
var StringZero = String("")

// String is a string.
type String string

// Concat returns the concatenation of s and other.
func (s String) Concat(other String) String {
	return s + other
}

// Equals implements Equatable.
func (s String) Equals(e Equatable) bool {
	return s == e.(String)
}

// Greater implements Comparable.
func (s String) Greater(c Comparable) bool {
	return s > c.(String)
}

// GreaterEqual implements Comparable.
func (s String) GreaterEqual(c Comparable) bool {
	return s >= c.(String)
}

// Less implements Comparable.
func (s String) Less(c Comparable) bool {
	return s < c.(String)
}

// LessEqual implements Comparable.
func (s String) LessEqual(c Comparable) bool {
	return s <= c.(String)
}

// NotEquals implements Equatable.
func (s String) NotEquals(e Equatable) bool {
	return s != e.(String)
}
