package goo

var _ Integer = Uint16(0)

// Uint16Zero is the Uint16 zero value.
var Uint16Zero = Uint16(0)

// Uint16 is a uint16.
type Uint16 uint16

// Add implements Number.
func (u Uint16) Add(n Number) Number {
	return u + n.(Uint16)
}

// And returns the bitwise conjunction of u and other.
func (u Uint16) And(other Integer) Integer {
	return u & other.(Uint16)
}

// Divide implements Number.
func (u Uint16) Divide(n Number) Number {
	return u / n.(Uint16)
}

// Equals implements Equatable.
func (u Uint16) Equals(e Equatable) bool {
	return u == e.(Uint16)
}

// Greater implements Comparable.
func (u Uint16) Greater(c Comparable) bool {
	return u > c.(Uint16)
}

// GreaterEqual implements Comparable.
func (u Uint16) GreaterEqual(c Comparable) bool {
	return u >= c.(Uint16)
}

// Less implements Comparable.
func (u Uint16) Less(c Comparable) bool {
	return u < c.(Uint16)
}

// LessEqual implements Comparable.
func (u Uint16) LessEqual(c Comparable) bool {
	return u <= c.(Uint16)
}

// Modulo implements Number.
func (u Uint16) Modulo(other Integer) Integer {
	return u % other.(Uint16)
}

// Multiply implements Number.
func (u Uint16) Multiply(n Number) Number {
	return u * n.(Uint16)
}

// Negate implements Number.
func (u Uint16) Negate() Number {
	return -u
}

// NotEquals implements Equatable.
func (u Uint16) NotEquals(e Equatable) bool {
	return u != e.(Uint16)
}

// Subtract implements Number.
func (u Uint16) Subtract(n Number) Number {
	return u - n.(Uint16)
}
