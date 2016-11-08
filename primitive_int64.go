package goo

// Int64Zero is the Int64 zero value.
var Int64Zero = Int64(0)

var _ Integer = Int64Zero

// Int64 is a int64.
type Int64 int64

// Add implements Number.
func (i Int64) Add(n Number) Number {
	return i + n.(Int64)
}

// And implements Integer.
func (i Int64) And(other Integer) Integer {
	return i & other.(Int64)
}

// AndNot implements Integer.
func (i Int64) AndNot(other Integer) Integer {
	return i &^ other.(Int64)
}

// Divide implements Number.
func (i Int64) Divide(other Number) Number {
	return i / other.(Int64)
}

// Equals implements Equatable.
func (i Int64) Equals(other Equatable) bool {
	return i == other.(Int64)
}

// Greater implements Comparable.
func (i Int64) Greater(other Comparable) bool {
	return i > other.(Int64)
}

// GreaterEqual implements Comparable.
func (i Int64) GreaterEqual(other Comparable) bool {
	return i >= other.(Int64)
}

// Left implements Integer.
func (i Int64) Left(n uint) Integer {
	return i << n
}

// Less implements Comparable.
func (i Int64) Less(other Comparable) bool {
	return i < other.(Int64)
}

// LessEqual implements Comparable.
func (i Int64) LessEqual(other Comparable) bool {
	return i <= other.(Int64)
}

// Modulo implements Integer.
func (i Int64) Modulo(other Integer) Integer {
	return i % other.(Int64)
}

// Multiply implements Number.
func (i Int64) Multiply(other Number) Number {
	return i * other.(Int64)
}

// Negate implements Number.
func (i Int64) Negate() Number {
	return -i
}

// Not implements Integer.
func (i Int64) Not() Integer {
	return ^i
}

// NotEquals implements Equatable.
func (i Int64) NotEquals(other Equatable) bool {
	return i != other.(Int64)
}

// Or implements Integer.
func (i Int64) Or(other Integer) Integer {
	return i | other.(Int64)
}

// Right implements Integer.
func (i Int64) Right(n uint) Integer {
	return i >> n
}

// Subtract implements Number.
func (i Int64) Subtract(other Number) Number {
	return i - other.(Int64)
}

// Xor implements Integer.
func (i Int64) Xor(other Integer) Integer {
	return i ^ other.(Int64)
}
