package goo

var _ Integer = Int8(0)

// Int8Zero is the Int8 zero value.
var Int8Zero = Int8(0)

// Int8 is a int8.
type Int8 int8

// Add implements Number.
func (i Int8) Add(n Number) Number {
	return i + n.(Int8)
}

// And returns the bitwise conjunction of i and other.
func (i Int8) And(other Integer) Integer {
	return i & other.(Int8)
}

// Divide implements Number.
func (i Int8) Divide(n Number) Number {
	return i / n.(Int8)
}

// Equals implements Equatable.
func (i Int8) Equals(e Equatable) bool {
	return i == e.(Int8)
}

// Greater implements Comparable.
func (i Int8) Greater(c Comparable) bool {
	return i > c.(Int8)
}

// GreaterEqual implements Comparable.
func (i Int8) GreaterEqual(c Comparable) bool {
	return i >= c.(Int8)
}

// Less implements Comparable.
func (i Int8) Less(c Comparable) bool {
	return i < c.(Int8)
}

// LessEqual implements Comparable.
func (i Int8) LessEqual(c Comparable) bool {
	return i <= c.(Int8)
}

// Modulo implements Number.
func (i Int8) Modulo(other Integer) Integer {
	return i % other.(Int8)
}

// Multiply implements Number.
func (i Int8) Multiply(n Number) Number {
	return i * n.(Int8)
}

// Negate implements Number.
func (i Int8) Negate() Number {
	return -i
}

// NotEquals implements Equatable.
func (i Int8) NotEquals(e Equatable) bool {
	return i != e.(Int8)
}

// Subtract implements Number.
func (i Int8) Subtract(n Number) Number {
	return i - n.(Int8)
}
