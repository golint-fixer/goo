package goo

// Uint16Zero is the Uint16 zero value.
var Uint16Zero = Uint16(0)

var _ Integer = Uint16Zero

// Uint16 is a uint16.
type Uint16 uint16

// Add implements Number.
func (u Uint16) Add(n Number) Number {
	return u + n.(Uint16)
}

// And implements Integer.
func (u Uint16) And(other Integer) Integer {
	return u & other.(Uint16)
}

// AndNot implements Integer.
func (u Uint16) AndNot(other Integer) Integer {
	return u &^ other.(Uint16)
}

// Divide implements Number.
func (u Uint16) Divide(other Number) Number {
	return u / other.(Uint16)
}

// Equals implements Equatable.
func (u Uint16) Equals(other Equatable) bool {
	return u == other.(Uint16)
}

// Greater implements Comparable.
func (u Uint16) Greater(other Comparable) bool {
	return u > other.(Uint16)
}

// GreaterEqual implements Comparable.
func (u Uint16) GreaterEqual(other Comparable) bool {
	return u >= other.(Uint16)
}

// Left implements Integer.
func (u Uint16) Left(n uint) Integer {
	return u << n
}

// Less implements Comparable.
func (u Uint16) Less(other Comparable) bool {
	return u < other.(Uint16)
}

// LessEqual implements Comparable.
func (u Uint16) LessEqual(other Comparable) bool {
	return u <= other.(Uint16)
}

// Modulo implements Integer.
func (u Uint16) Modulo(other Integer) Integer {
	return u % other.(Uint16)
}

// Multiply implements Number.
func (u Uint16) Multiply(other Number) Number {
	return u * other.(Uint16)
}

// Negate implements Number.
func (u Uint16) Negate() Number {
	return -u
}

// Not implements Integer.
func (u Uint16) Not() Integer {
	return ^u
}

// NotEquals implements Equatable.
func (u Uint16) NotEquals(other Equatable) bool {
	return u != other.(Uint16)
}

// Or implements Integer.
func (u Uint16) Or(other Integer) Integer {
	return u | other.(Uint16)
}

// Right implements Integer.
func (u Uint16) Right(n uint) Integer {
	return u >> n
}

// Subtract implements Number.
func (u Uint16) Subtract(other Number) Number {
	return u - other.(Uint16)
}

// Xor implements Integer.
func (u Uint16) Xor(other Integer) Integer {
	return u ^ other.(Uint16)
}
