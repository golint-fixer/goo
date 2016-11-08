package goo

// UintZero is the Uint zero value.
var UintZero = Uint(0)

var _ Integer = UintZero

var _ Pointer = &UintZero

var _ Value = UintZero

// Uint is a uint.
type Uint uint

// Add implements Number.
func (u Uint) Add(n Number) Number {
	return u + n.(Uint)
}

// And implements Integer.
func (u Uint) And(other Integer) Integer {
	return u & other.(Uint)
}

// AndNot implements Integer.
func (u Uint) AndNot(other Integer) Integer {
	return u &^ other.(Uint)
}

// Dereference implements Pointer.
func (u *Uint) Dereference() Value {
	return *u
}

// Divide implements Number.
func (u Uint) Divide(other Number) Number {
	return u / other.(Uint)
}

// Equals implements Equatable.
func (u Uint) Equals(other Equatable) bool {
	return u == other.(Uint)
}

// Greater implements Comparable.
func (u Uint) Greater(other Comparable) bool {
	return u > other.(Uint)
}

// GreaterEqual implements Comparable.
func (u Uint) GreaterEqual(other Comparable) bool {
	return u >= other.(Uint)
}

// Left implements Integer.
func (u Uint) Left(n uint) Integer {
	return u << n
}

// Less implements Comparable.
func (u Uint) Less(other Comparable) bool {
	return u < other.(Uint)
}

// LessEqual implements Comparable.
func (u Uint) LessEqual(other Comparable) bool {
	return u <= other.(Uint)
}

// Modulo implements Integer.
func (u Uint) Modulo(other Integer) Integer {
	return u % other.(Uint)
}

// Multiply implements Number.
func (u Uint) Multiply(other Number) Number {
	return u * other.(Uint)
}

// Negate implements Number.
func (u Uint) Negate() Number {
	return -u
}

// Not implements Integer.
func (u Uint) Not() Integer {
	return ^u
}

// NotEquals implements Equatable.
func (u Uint) NotEquals(other Equatable) bool {
	return u != other.(Uint)
}

// Or implements Integer.
func (u Uint) Or(other Integer) Integer {
	return u | other.(Uint)
}

// Reference implements Value.
func (u Uint) Reference() Pointer {
	return &u
}

// Right implements Integer.
func (u Uint) Right(n uint) Integer {
	return u >> n
}

// Subtract implements Number.
func (u Uint) Subtract(other Number) Number {
	return u - other.(Uint)
}

// Xor implements Integer.
func (u Uint) Xor(other Integer) Integer {
	return u ^ other.(Uint)
}
