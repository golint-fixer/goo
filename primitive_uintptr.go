package goo

var _ Integer = Uintptr(0)

// UintptrZero is the Uintptr zero value.
var UintptrZero = Uintptr(0)

// Uintptr is a uintptr.
type Uintptr uintptr

// Add implements Number.
func (u Uintptr) Add(n Number) Number {
	return u + n.(Uintptr)
}

// And returns the bitwise conjunction of u and other.
func (u Uintptr) And(other Integer) Integer {
	return u & other.(Uintptr)
}

// Divide implements Number.
func (u Uintptr) Divide(n Number) Number {
	return u / n.(Uintptr)
}

// Equals implements Equatable.
func (u Uintptr) Equals(e Equatable) bool {
	return u == e.(Uintptr)
}

// Greater implements Comparable.
func (u Uintptr) Greater(c Comparable) bool {
	return u > c.(Uintptr)
}

// GreaterEqual implements Comparable.
func (u Uintptr) GreaterEqual(c Comparable) bool {
	return u >= c.(Uintptr)
}

// Less implements Comparable.
func (u Uintptr) Less(c Comparable) bool {
	return u < c.(Uintptr)
}

// LessEqual implements Comparable.
func (u Uintptr) LessEqual(c Comparable) bool {
	return u <= c.(Uintptr)
}

// Modulo implements Number.
func (u Uintptr) Modulo(other Integer) Integer {
	return u % other.(Uintptr)
}

// Multiply implements Number.
func (u Uintptr) Multiply(n Number) Number {
	return u * n.(Uintptr)
}

// Negate implements Number.
func (u Uintptr) Negate() Number {
	return -u
}

// NotEquals implements Equatable.
func (u Uintptr) NotEquals(e Equatable) bool {
	return u != e.(Uintptr)
}

// Subtract implements Number.
func (u Uintptr) Subtract(n Number) Number {
	return u - n.(Uintptr)
}
