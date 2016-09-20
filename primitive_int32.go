package goo

var _ Integer = Int32(0)

// Int32Zero is the Int32 zero value.
var Int32Zero = Int32(0)

// Int32 is a int32.
type Int32 int32

// Add implements Number.
func (i Int32) Add(n Number) Number {
	return i + n.(Int32)
}

// And returns the bitwise conjunction of i and other.
func (i Int32) And(other Integer) Integer {
	return i & other.(Int32)
}

// Divide implements Number.
func (i Int32) Divide(n Number) Number {
	return i / n.(Int32)
}

// Equals implements Equatable.
func (i Int32) Equals(e Equatable) bool {
	return i == e.(Int32)
}

// Greater implements Comparable.
func (i Int32) Greater(c Comparable) bool {
	return i > c.(Int32)
}

// GreaterEqual implements Comparable.
func (i Int32) GreaterEqual(c Comparable) bool {
	return i >= c.(Int32)
}

// Less implements Comparable.
func (i Int32) Less(c Comparable) bool {
	return i < c.(Int32)
}

// LessEqual implements Comparable.
func (i Int32) LessEqual(c Comparable) bool {
	return i <= c.(Int32)
}

// Modulo implements Number.
func (i Int32) Modulo(other Integer) Integer {
	return i % other.(Int32)
}

// Multiply implements Number.
func (i Int32) Multiply(n Number) Number {
	return i * n.(Int32)
}

// Negate implements Number.
func (i Int32) Negate() Number {
	return -i
}

// NotEquals implements Equatable.
func (i Int32) NotEquals(e Equatable) bool {
	return i != e.(Int32)
}

// Subtract implements Number.
func (i Int32) Subtract(n Number) Number {
	return i - n.(Int32)
}
