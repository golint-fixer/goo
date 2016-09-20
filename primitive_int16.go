package goo

var _ Integer = Int16(0)

// Int16Zero is the Int16 zero value.
var Int16Zero = Int16(0)

// Int16 is a int16.
type Int16 int16

// Add implements Number.
func (i Int16) Add(n Number) Number {
	return i + n.(Int16)
}

// And returns the bitwise conjunction of i and other.
func (i Int16) And(other Integer) Integer {
	return i & other.(Int16)
}

// Divide implements Number.
func (i Int16) Divide(n Number) Number {
	return i / n.(Int16)
}

// Equals implements Equatable.
func (i Int16) Equals(e Equatable) bool {
	return i == e.(Int16)
}

// Greater implements Comparable.
func (i Int16) Greater(c Comparable) bool {
	return i > c.(Int16)
}

// GreaterEqual implements Comparable.
func (i Int16) GreaterEqual(c Comparable) bool {
	return i >= c.(Int16)
}

// Less implements Comparable.
func (i Int16) Less(c Comparable) bool {
	return i < c.(Int16)
}

// LessEqual implements Comparable.
func (i Int16) LessEqual(c Comparable) bool {
	return i <= c.(Int16)
}

// Modulo implements Number.
func (i Int16) Modulo(other Integer) Integer {
	return i % other.(Int16)
}

// Multiply implements Number.
func (i Int16) Multiply(n Number) Number {
	return i * n.(Int16)
}

// Negate implements Number.
func (i Int16) Negate() Number {
	return -i
}

// NotEquals implements Equatable.
func (i Int16) NotEquals(e Equatable) bool {
	return i != e.(Int16)
}

// Subtract implements Number.
func (i Int16) Subtract(n Number) Number {
	return i - n.(Int16)
}
