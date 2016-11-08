package goo

// Int8Zero is the Int8 zero value.
var Int8Zero = Int8(0)

var _ Integer = Int8Zero

var _ Pointer = &Int8Zero

var _ Value = Int8Zero

// Int8 is a int8.
type Int8 int8

// Add implements Number.
func (i Int8) Add(n Number) Number {
	return i + n.(Int8)
}

// And implements Integer.
func (i Int8) And(other Integer) Integer {
	return i & other.(Int8)
}

// AndNot implements Integer.
func (i Int8) AndNot(other Integer) Integer {
	return i &^ other.(Int8)
}

// Dereference implements Pointer.
func (i *Int8) Dereference() Value {
	return *i
}

// Divide implements Number.
func (i Int8) Divide(other Number) Number {
	return i / other.(Int8)
}

// Equals implements Equatable.
func (i Int8) Equals(other Equatable) bool {
	return i == other.(Int8)
}

// Greater implements Comparable.
func (i Int8) Greater(other Comparable) bool {
	return i > other.(Int8)
}

// GreaterEqual implements Comparable.
func (i Int8) GreaterEqual(other Comparable) bool {
	return i >= other.(Int8)
}

// Left implements Integer.
func (i Int8) Left(n uint) Integer {
	return i << n
}

// Less implements Comparable.
func (i Int8) Less(other Comparable) bool {
	return i < other.(Int8)
}

// LessEqual implements Comparable.
func (i Int8) LessEqual(other Comparable) bool {
	return i <= other.(Int8)
}

// Modulo implements Integer.
func (i Int8) Modulo(other Integer) Integer {
	return i % other.(Int8)
}

// Multiply implements Number.
func (i Int8) Multiply(other Number) Number {
	return i * other.(Int8)
}

// Negate implements Number.
func (i Int8) Negate() Number {
	return -i
}

// Not implements Integer.
func (i Int8) Not() Integer {
	return ^i
}

// NotEquals implements Equatable.
func (i Int8) NotEquals(other Equatable) bool {
	return i != other.(Int8)
}

// Or implements Integer.
func (i Int8) Or(other Integer) Integer {
	return i | other.(Int8)
}

// Reference implements Value.
func (i Int8) Reference() Pointer {
	return &i
}

// Right implements Integer.
func (i Int8) Right(n uint) Integer {
	return i >> n
}

// Subtract implements Number.
func (i Int8) Subtract(other Number) Number {
	return i - other.(Int8)
}

// Xor implements Integer.
func (i Int8) Xor(other Integer) Integer {
	return i ^ other.(Int8)
}
