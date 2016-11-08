package goo

// Uint64Zero is the Uint64 zero value.
var Uint64Zero = Uint64(0)

var _ Integer = Uint64Zero

var _ Pointer = &Uint64Zero

var _ Value = Uint64Zero

// Uint64 is a uint64.
type Uint64 uint64

// Add implements Number.
func (u Uint64) Add(n Number) Number {
	return u + n.(Uint64)
}

// And implements Integer.
func (u Uint64) And(other Integer) Integer {
	return u & other.(Uint64)
}

// AndNot implements Integer.
func (u Uint64) AndNot(other Integer) Integer {
	return u &^ other.(Uint64)
}

// Dereference implements Pointer.
func (u *Uint64) Dereference() Value {
	return *u
}

// Divide implements Number.
func (u Uint64) Divide(other Number) Number {
	return u / other.(Uint64)
}

// Equals implements Equatable.
func (u Uint64) Equals(other Equatable) bool {
	return u == other.(Uint64)
}

// Greater implements Comparable.
func (u Uint64) Greater(other Comparable) bool {
	return u > other.(Uint64)
}

// GreaterEqual implements Comparable.
func (u Uint64) GreaterEqual(other Comparable) bool {
	return u >= other.(Uint64)
}

// Left implements Integer.
func (u Uint64) Left(n uint) Integer {
	return u << n
}

// Less implements Comparable.
func (u Uint64) Less(other Comparable) bool {
	return u < other.(Uint64)
}

// LessEqual implements Comparable.
func (u Uint64) LessEqual(other Comparable) bool {
	return u <= other.(Uint64)
}

// Modulo implements Integer.
func (u Uint64) Modulo(other Integer) Integer {
	return u % other.(Uint64)
}

// Multiply implements Number.
func (u Uint64) Multiply(other Number) Number {
	return u * other.(Uint64)
}

// Negate implements Number.
func (u Uint64) Negate() Number {
	return -u
}

// Not implements Integer.
func (u Uint64) Not() Integer {
	return ^u
}

// NotEquals implements Equatable.
func (u Uint64) NotEquals(other Equatable) bool {
	return u != other.(Uint64)
}

// Or implements Integer.
func (u Uint64) Or(other Integer) Integer {
	return u | other.(Uint64)
}

// Reference implements Value.
func (u Uint64) Reference() Pointer {
	return &u
}

// Right implements Integer.
func (u Uint64) Right(n uint) Integer {
	return u >> n
}

// Subtract implements Number.
func (u Uint64) Subtract(other Number) Number {
	return u - other.(Uint64)
}

// Xor implements Integer.
func (u Uint64) Xor(other Integer) Integer {
	return u ^ other.(Uint64)
}
