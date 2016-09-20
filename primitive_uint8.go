package goo

var _ Integer = Uint8(0)

// Uint8Zero is the Uint8 zero value.
var Uint8Zero = Uint8(0)

// Uint8 is a uint8.
type Uint8 uint8

// Add implements Number.
func (u Uint8) Add(n Number) Number {
	return u + n.(Uint8)
}

// And returns the bitwise conjunction of u and other.
func (u Uint8) And(other Integer) Integer {
	return u & other.(Uint8)
}

// Divide implements Number.
func (u Uint8) Divide(n Number) Number {
	return u / n.(Uint8)
}

// Equals implements Equatable.
func (u Uint8) Equals(e Equatable) bool {
	return u == e.(Uint8)
}

// Greater implements Comparable.
func (u Uint8) Greater(c Comparable) bool {
	return u > c.(Uint8)
}

// GreaterEqual implements Comparable.
func (u Uint8) GreaterEqual(c Comparable) bool {
	return u >= c.(Uint8)
}

// Less implements Comparable.
func (u Uint8) Less(c Comparable) bool {
	return u < c.(Uint8)
}

// LessEqual implements Comparable.
func (u Uint8) LessEqual(c Comparable) bool {
	return u <= c.(Uint8)
}

// Modulo implements Number.
func (u Uint8) Modulo(other Integer) Integer {
	return u % other.(Uint8)
}

// Multiply implements Number.
func (u Uint8) Multiply(n Number) Number {
	return u * n.(Uint8)
}

// Negate implements Number.
func (u Uint8) Negate() Number {
	return -u
}

// NotEquals implements Equatable.
func (u Uint8) NotEquals(e Equatable) bool {
	return u != e.(Uint8)
}

// Subtract implements Number.
func (u Uint8) Subtract(n Number) Number {
	return u - n.(Uint8)
}
