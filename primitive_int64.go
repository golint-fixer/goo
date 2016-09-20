package goo

var _ Integer = Int64(0)

// Int64Zero is the Int64 zero value.
var Int64Zero = Int64(0)

// Int64 is a int64.
type Int64 int64

// Add implements Number.
func (i Int64) Add(n Number) Number {
	return i + n.(Int64)
}

// And returns the bitwise conjunction of i and other.
func (i Int64) And(other Integer) Integer {
	return i & other.(Int64)
}

// Divide implements Number.
func (i Int64) Divide(n Number) Number {
	return i / n.(Int64)
}

// Equals implements Equatable.
func (i Int64) Equals(e Equatable) bool {
	return i == e.(Int64)
}

// Greater implements Comparable.
func (i Int64) Greater(c Comparable) bool {
	return i > c.(Int64)
}

// GreaterEqual implements Comparable.
func (i Int64) GreaterEqual(c Comparable) bool {
	return i >= c.(Int64)
}

// Less implements Comparable.
func (i Int64) Less(c Comparable) bool {
	return i < c.(Int64)
}

// LessEqual implements Comparable.
func (i Int64) LessEqual(c Comparable) bool {
	return i <= c.(Int64)
}

// Modulo implements Number.
func (i Int64) Modulo(other Integer) Integer {
	return i % other.(Int64)
}

// Multiply implements Number.
func (i Int64) Multiply(n Number) Number {
	return i * n.(Int64)
}

// Negate implements Number.
func (i Int64) Negate() Number {
	return -i
}

// NotEquals implements Equatable.
func (i Int64) NotEquals(e Equatable) bool {
	return i != e.(Int64)
}

// Subtract implements Number.
func (i Int64) Subtract(n Number) Number {
	return i - n.(Int64)
}
