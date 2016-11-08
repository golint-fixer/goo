package goo

// StringZero is the String zero value.
var StringZero = String("")

var _ Comparable = StringZero

// String is a string.
type String string

// Concat returns the concatenation of s and other.
func (s String) Concat(other String) String {
	return s + other
}

// Equals implements Equatable.
func (s String) Equals(other Equatable) bool {
	return s == other.(String)
}

// Greater implements Comparable.
func (s String) Greater(other Comparable) bool {
	return s > other.(String)
}

// GreaterEqual implements Comparable.
func (s String) GreaterEqual(other Comparable) bool {
	return s >= other.(String)
}

// Less implements Comparable.
func (s String) Less(other Comparable) bool {
	return s < other.(String)
}

// LessEqual implements Comparable.
func (s String) LessEqual(other Comparable) bool {
	return s <= other.(String)
}

// NotEquals implements Equatable.
func (s String) NotEquals(other Equatable) bool {
	return s != other.(String)
}
