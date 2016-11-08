package goo

// Int32Zero is the Int32 zero value.
var Int32Zero = Int32(0)

var _ Integer = Int32Zero

// Int32 is a int32.
type Int32 int32

// Add implements Number.
func (i Int32) Add(n Number) Number {
	return i + n.(Int32)
}

// And implements Integer.
func (i Int32) And(other Integer) Integer {
	return i & other.(Int32)
}

// AndNot implements Integer.
func (i Int32) AndNot(other Integer) Integer {
	return i &^ other.(Int32)
}

// Divide implements Number.
func (i Int32) Divide(other Number) Number {
	return i / other.(Int32)
}

// Equals implements Equatable.
func (i Int32) Equals(other Equatable) bool {
	return i == other.(Int32)
}

// Greater implements Comparable.
func (i Int32) Greater(other Comparable) bool {
	return i > other.(Int32)
}

// GreaterEqual implements Comparable.
func (i Int32) GreaterEqual(other Comparable) bool {
	return i >= other.(Int32)
}

// Left implements Integer.
func (i Int32) Left(n uint) Integer {
	return i << n
}

// Less implements Comparable.
func (i Int32) Less(other Comparable) bool {
	return i < other.(Int32)
}

// LessEqual implements Comparable.
func (i Int32) LessEqual(other Comparable) bool {
	return i <= other.(Int32)
}

// Modulo implements Integer.
func (i Int32) Modulo(other Integer) Integer {
	return i % other.(Int32)
}

// Multiply implements Number.
func (i Int32) Multiply(other Number) Number {
	return i * other.(Int32)
}

// Negate implements Number.
func (i Int32) Negate() Number {
	return -i
}

// Not implements Integer.
func (i Int32) Not() Integer {
	return ^i
}

// NotEquals implements Equatable.
func (i Int32) NotEquals(other Equatable) bool {
	return i != other.(Int32)
}

// Or implements Integer.
func (i Int32) Or(other Integer) Integer {
	return i | other.(Int32)
}

// Right implements Integer.
func (i Int32) Right(n uint) Integer {
	return i >> n
}

// Subtract implements Number.
func (i Int32) Subtract(other Number) Number {
	return i - other.(Int32)
}

// Xor implements Integer.
func (i Int32) Xor(other Integer) Integer {
	return i ^ other.(Int32)
}
