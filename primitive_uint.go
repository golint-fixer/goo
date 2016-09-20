package goo

var _ Integer = Uint(0)

// UintZero is the Uint zero value.
var UintZero = Uint(0)

// Uint is a uint.
type Uint uint

// Add implements Number.
func (u Uint) Add(n Number) Number {
	return u + n.(Uint)
}

// And returns the bitwise conjunction of u and other.
func (u Uint) And(other Integer) Integer {
	return u & other.(Uint)
}

// Divide implements Number.
func (u Uint) Divide(n Number) Number {
	return u / n.(Uint)
}

// Equals implements Equatable.
func (u Uint) Equals(e Equatable) bool {
	return u == e.(Uint)
}

// Greater implements Comparable.
func (u Uint) Greater(c Comparable) bool {
	return u > c.(Uint)
}

// GreaterEqual implements Comparable.
func (u Uint) GreaterEqual(c Comparable) bool {
	return u >= c.(Uint)
}

// Less implements Comparable.
func (u Uint) Less(c Comparable) bool {
	return u < c.(Uint)
}

// LessEqual implements Comparable.
func (u Uint) LessEqual(c Comparable) bool {
	return u <= c.(Uint)
}

// Modulo implements Number.
func (u Uint) Modulo(other Integer) Integer {
	return u % other.(Uint)
}

// Multiply implements Number.
func (u Uint) Multiply(n Number) Number {
	return u * n.(Uint)
}

// Negate implements Number.
func (u Uint) Negate() Number {
	return -u
}

// NotEquals implements Equatable.
func (u Uint) NotEquals(e Equatable) bool {
	return u != e.(Uint)
}

// Subtract implements Number.
func (u Uint) Subtract(n Number) Number {
	return u - n.(Uint)
}
