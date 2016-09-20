package goo

var _ Integer = Int(0)

// IntZero is the Int zero value.
var IntZero = Int(0)

// Int is a int.
type Int int

// Add implements Number.
func (i Int) Add(n Number) Number {
	return i + n.(Int)
}

// And returns the bitwise conjunction of i and other.
func (i Int) And(other Integer) Integer {
	return i & other.(Int)
}

// Divide implements Number.
func (i Int) Divide(n Number) Number {
	return i / n.(Int)
}

// Equals implements Equatable.
func (i Int) Equals(e Equatable) bool {
	return i == e.(Int)
}

// Greater implements Comparable.
func (i Int) Greater(c Comparable) bool {
	return i > c.(Int)
}

// GreaterEqual implements Comparable.
func (i Int) GreaterEqual(c Comparable) bool {
	return i >= c.(Int)
}

// Less implements Comparable.
func (i Int) Less(c Comparable) bool {
	return i < c.(Int)
}

// LessEqual implements Comparable.
func (i Int) LessEqual(c Comparable) bool {
	return i <= c.(Int)
}

// Modulo implements Number.
func (i Int) Modulo(other Integer) Integer {
	return i % other.(Int)
}

// Multiply implements Number.
func (i Int) Multiply(n Number) Number {
	return i * n.(Int)
}

// Negate implements Number.
func (i Int) Negate() Number {
	return -i
}

// NotEquals implements Equatable.
func (i Int) NotEquals(e Equatable) bool {
	return i != e.(Int)
}

// Subtract implements Number.
func (i Int) Subtract(n Number) Number {
	return i - n.(Int)
}
