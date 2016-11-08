package goo

// Uint8Zero is the Uint8 zero value.
var Uint8Zero = Uint8(0)

var _ Integer = Uint8Zero

// Uint8 is a uint8.
type Uint8 uint8

// Add implements Number.
func (u Uint8) Add(n Number) Number {
	return u + n.(Uint8)
}

// And implements Integer.
func (u Uint8) And(other Integer) Integer {
	return u & other.(Uint8)
}

// AndNot implements Integer.
func (u Uint8) AndNot(other Integer) Integer {
	return u &^ other.(Uint8)
}

// Divide implements Number.
func (u Uint8) Divide(other Number) Number {
	return u / other.(Uint8)
}

// Equals implements Equatable.
func (u Uint8) Equals(other Equatable) bool {
	return u == other.(Uint8)
}

// Greater implements Comparable.
func (u Uint8) Greater(other Comparable) bool {
	return u > other.(Uint8)
}

// GreaterEqual implements Comparable.
func (u Uint8) GreaterEqual(other Comparable) bool {
	return u >= other.(Uint8)
}

// Left implements Integer.
func (u Uint8) Left(n uint) Integer {
	return u << n
}

// Less implements Comparable.
func (u Uint8) Less(other Comparable) bool {
	return u < other.(Uint8)
}

// LessEqual implements Comparable.
func (u Uint8) LessEqual(other Comparable) bool {
	return u <= other.(Uint8)
}

// Modulo implements Integer.
func (u Uint8) Modulo(other Integer) Integer {
	return u % other.(Uint8)
}

// Multiply implements Number.
func (u Uint8) Multiply(other Number) Number {
	return u * other.(Uint8)
}

// Negate implements Number.
func (u Uint8) Negate() Number {
	return -u
}

// Not implements Integer.
func (u Uint8) Not() Integer {
	return ^u
}

// NotEquals implements Equatable.
func (u Uint8) NotEquals(other Equatable) bool {
	return u != other.(Uint8)
}

// Or implements Integer.
func (u Uint8) Or(other Integer) Integer {
	return u | other.(Uint8)
}

// Right implements Integer.
func (u Uint8) Right(n uint) Integer {
	return u >> n
}

// Subtract implements Number.
func (u Uint8) Subtract(other Number) Number {
	return u - other.(Uint8)
}

// Xor implements Integer.
func (u Uint8) Xor(other Integer) Integer {
	return u ^ other.(Uint8)
}
