package goo

var _ Integer = Byte(0)

// ByteZero is the Byte zero value.
var ByteZero = Byte(0)

// Byte is a byte.
type Byte byte

// Add implements Number.
func (b Byte) Add(n Number) Number {
	return b + n.(Byte)
}

// And returns the bitwise conjunction of b and other.
func (b Byte) And(other Integer) Integer {
	return b & other.(Byte)
}

// Divide implements Number.
func (b Byte) Divide(n Number) Number {
	return b / n.(Byte)
}

// Equals implements Equatable.
func (b Byte) Equals(e Equatable) bool {
	return b == e.(Byte)
}

// Greater implements Comparable.
func (b Byte) Greater(c Comparable) bool {
	return b > c.(Byte)
}

// GreaterEqual implements Comparable.
func (b Byte) GreaterEqual(c Comparable) bool {
	return b >= c.(Byte)
}

// Less implements Comparable.
func (b Byte) Less(c Comparable) bool {
	return b < c.(Byte)
}

// LessEqual implements Comparable.
func (b Byte) LessEqual(c Comparable) bool {
	return b <= c.(Byte)
}

// Modulo implements Number.
func (b Byte) Modulo(other Integer) Integer {
	return b % other.(Byte)
}

// Multiply implements Number.
func (b Byte) Multiply(n Number) Number {
	return b * n.(Byte)
}

// Negate implements Number.
func (b Byte) Negate() Number {
	return -b
}

// NotEquals implements Equatable.
func (b Byte) NotEquals(e Equatable) bool {
	return b != e.(Byte)
}

// Subtract implements Number.
func (b Byte) Subtract(n Number) Number {
	return b - n.(Byte)
}
