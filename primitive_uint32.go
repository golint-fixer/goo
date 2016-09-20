package goo

var _ Integer = Uint32(0)

// Uint32Zero is the Uint32 zero value.
var Uint32Zero = Uint32(0)

// Uint32 is a uint32.
type Uint32 uint32

// Add implements Number.
func (u Uint32) Add(n Number) Number {
	return u + n.(Uint32)
}

// And returns the bitwise conjunction of u and other.
func (u Uint32) And(other Integer) Integer {
	return u & other.(Uint32)
}

// Divide implements Number.
func (u Uint32) Divide(n Number) Number {
	return u / n.(Uint32)
}

// Equals implements Equatable.
func (u Uint32) Equals(e Equatable) bool {
	return u == e.(Uint32)
}

// Greater implements Comparable.
func (u Uint32) Greater(c Comparable) bool {
	return u > c.(Uint32)
}

// GreaterEqual implements Comparable.
func (u Uint32) GreaterEqual(c Comparable) bool {
	return u >= c.(Uint32)
}

// Less implements Comparable.
func (u Uint32) Less(c Comparable) bool {
	return u < c.(Uint32)
}

// LessEqual implements Comparable.
func (u Uint32) LessEqual(c Comparable) bool {
	return u <= c.(Uint32)
}

// Modulo implements Number.
func (u Uint32) Modulo(other Integer) Integer {
	return u % other.(Uint32)
}

// Multiply implements Number.
func (u Uint32) Multiply(n Number) Number {
	return u * n.(Uint32)
}

// Negate implements Number.
func (u Uint32) Negate() Number {
	return -u
}

// NotEquals implements Equatable.
func (u Uint32) NotEquals(e Equatable) bool {
	return u != e.(Uint32)
}

// Subtract implements Number.
func (u Uint32) Subtract(n Number) Number {
	return u - n.(Uint32)
}
