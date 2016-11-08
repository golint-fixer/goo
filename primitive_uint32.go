package goo

// Uint32Zero is the Uint32 zero value.
var Uint32Zero = Uint32(0)

var _ Integer = Uint32Zero

var _ Pointer = &Uint32Zero

var _ Value = Uint32Zero

// Uint32 is a uint32.
type Uint32 uint32

// Add implements Number.
func (u Uint32) Add(n Number) Number {
	return u + n.(Uint32)
}

// And implements Integer.
func (u Uint32) And(other Integer) Integer {
	return u & other.(Uint32)
}

// AndNot implements Integer.
func (u Uint32) AndNot(other Integer) Integer {
	return u &^ other.(Uint32)
}

// Dereference implements Pointer.
func (u *Uint32) Dereference() Value {
	return *u
}

// Divide implements Number.
func (u Uint32) Divide(other Number) Number {
	return u / other.(Uint32)
}

// Equals implements Equatable.
func (u Uint32) Equals(other Equatable) bool {
	return u == other.(Uint32)
}

// Greater implements Comparable.
func (u Uint32) Greater(other Comparable) bool {
	return u > other.(Uint32)
}

// GreaterEqual implements Comparable.
func (u Uint32) GreaterEqual(other Comparable) bool {
	return u >= other.(Uint32)
}

// Left implements Integer.
func (u Uint32) Left(n uint) Integer {
	return u << n
}

// Less implements Comparable.
func (u Uint32) Less(other Comparable) bool {
	return u < other.(Uint32)
}

// LessEqual implements Comparable.
func (u Uint32) LessEqual(other Comparable) bool {
	return u <= other.(Uint32)
}

// Modulo implements Integer.
func (u Uint32) Modulo(other Integer) Integer {
	return u % other.(Uint32)
}

// Multiply implements Number.
func (u Uint32) Multiply(other Number) Number {
	return u * other.(Uint32)
}

// Negate implements Number.
func (u Uint32) Negate() Number {
	return -u
}

// Not implements Integer.
func (u Uint32) Not() Integer {
	return ^u
}

// NotEquals implements Equatable.
func (u Uint32) NotEquals(other Equatable) bool {
	return u != other.(Uint32)
}

// Or implements Integer.
func (u Uint32) Or(other Integer) Integer {
	return u | other.(Uint32)
}

// Reference implements Value.
func (u Uint32) Reference() Pointer {
	return &u
}

// Right implements Integer.
func (u Uint32) Right(n uint) Integer {
	return u >> n
}

// Subtract implements Number.
func (u Uint32) Subtract(other Number) Number {
	return u - other.(Uint32)
}

// Xor implements Integer.
func (u Uint32) Xor(other Integer) Integer {
	return u ^ other.(Uint32)
}
