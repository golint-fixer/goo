package goo

// UintptrZero is the Uintptr zero value.
var UintptrZero = Uintptr(0)

var _ Integer = UintptrZero

var _ Pointer = &UintptrZero

var _ Value = UintptrZero

// Uintptr is a uintptr.
type Uintptr uintptr

// Add implements Number.
func (u Uintptr) Add(n Number) Number {
	return u + n.(Uintptr)
}

// And implements Integer.
func (u Uintptr) And(other Integer) Integer {
	return u & other.(Uintptr)
}

// AndNot implements Integer.
func (u Uintptr) AndNot(other Integer) Integer {
	return u &^ other.(Uintptr)
}

// Dereference implements Pointer.
func (u *Uintptr) Dereference() Value {
	return *u
}

// Divide implements Number.
func (u Uintptr) Divide(other Number) Number {
	return u / other.(Uintptr)
}

// Equals implements Equatable.
func (u Uintptr) Equals(other Equatable) bool {
	return u == other.(Uintptr)
}

// Greater implements Comparable.
func (u Uintptr) Greater(other Comparable) bool {
	return u > other.(Uintptr)
}

// GreaterEqual implements Comparable.
func (u Uintptr) GreaterEqual(other Comparable) bool {
	return u >= other.(Uintptr)
}

// Left implements Integer.
func (u Uintptr) Left(n uint) Integer {
	return u << n
}

// Less implements Comparable.
func (u Uintptr) Less(other Comparable) bool {
	return u < other.(Uintptr)
}

// LessEqual implements Comparable.
func (u Uintptr) LessEqual(other Comparable) bool {
	return u <= other.(Uintptr)
}

// Modulo implements Integer.
func (u Uintptr) Modulo(other Integer) Integer {
	return u % other.(Uintptr)
}

// Multiply implements Number.
func (u Uintptr) Multiply(other Number) Number {
	return u * other.(Uintptr)
}

// Negate implements Number.
func (u Uintptr) Negate() Number {
	return -u
}

// Not implements Integer.
func (u Uintptr) Not() Integer {
	return ^u
}

// NotEquals implements Equatable.
func (u Uintptr) NotEquals(other Equatable) bool {
	return u != other.(Uintptr)
}

// Or implements Integer.
func (u Uintptr) Or(other Integer) Integer {
	return u | other.(Uintptr)
}

// Reference implements Value.
func (u Uintptr) Reference() Pointer {
	return &u
}

// Right implements Integer.
func (u Uintptr) Right(n uint) Integer {
	return u >> n
}

// Subtract implements Number.
func (u Uintptr) Subtract(other Number) Number {
	return u - other.(Uintptr)
}

// Xor implements Integer.
func (u Uintptr) Xor(other Integer) Integer {
	return u ^ other.(Uintptr)
}
