package goo

var _ Integer = Uint64(0)

// Uint64Zero is the Uint64 zero value.
var Uint64Zero = Uint64(0)

// Uint64 is a uint64.
type Uint64 uint64

// Add implements Number.
func (u Uint64) Add(n Number) Number {
	return u + n.(Uint64)
}

// And returns the bitwise conjunction of u and other.
func (u Uint64) And(other Integer) Integer {
	return u & other.(Uint64)
}

// Divide implements Number.
func (u Uint64) Divide(n Number) Number {
	return u / n.(Uint64)
}

// Equals implements Equatable.
func (u Uint64) Equals(e Equatable) bool {
	return u == e.(Uint64)
}

// Greater implements Comparable.
func (u Uint64) Greater(c Comparable) bool {
	return u > c.(Uint64)
}

// GreaterEqual implements Comparable.
func (u Uint64) GreaterEqual(c Comparable) bool {
	return u >= c.(Uint64)
}

// Less implements Comparable.
func (u Uint64) Less(c Comparable) bool {
	return u < c.(Uint64)
}

// LessEqual implements Comparable.
func (u Uint64) LessEqual(c Comparable) bool {
	return u <= c.(Uint64)
}

// Modulo implements Number.
func (u Uint64) Modulo(other Integer) Integer {
	return u % other.(Uint64)
}

// Multiply implements Number.
func (u Uint64) Multiply(n Number) Number {
	return u * n.(Uint64)
}

// Negate implements Number.
func (u Uint64) Negate() Number {
	return -u
}

// NotEquals implements Equatable.
func (u Uint64) NotEquals(e Equatable) bool {
	return u != e.(Uint64)
}

// Subtract implements Number.
func (u Uint64) Subtract(n Number) Number {
	return u - n.(Uint64)
}
