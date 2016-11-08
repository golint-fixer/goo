package goo

// Int16Zero is the Int16 zero value.
var Int16Zero = Int16(0)

var _ Integer = Int16Zero

var _ Pointer = &Int16Zero

var _ Value = Int16Zero

// Int16 is a int16.
type Int16 int16

// Add implements Number.
func (i Int16) Add(n Number) Number {
	return i + n.(Int16)
}

// And implements Integer.
func (i Int16) And(other Integer) Integer {
	return i & other.(Int16)
}

// AndNot implements Integer.
func (i Int16) AndNot(other Integer) Integer {
	return i &^ other.(Int16)
}

// Dereference implements Pointer.
func (i *Int16) Dereference() Value {
	return *i
}

// Divide implements Number.
func (i Int16) Divide(other Number) Number {
	return i / other.(Int16)
}

// Equals implements Equatable.
func (i Int16) Equals(other Equatable) bool {
	return i == other.(Int16)
}

// Greater implements Comparable.
func (i Int16) Greater(other Comparable) bool {
	return i > other.(Int16)
}

// GreaterEqual implements Comparable.
func (i Int16) GreaterEqual(other Comparable) bool {
	return i >= other.(Int16)
}

// Left implements Integer.
func (i Int16) Left(n uint) Integer {
	return i << n
}

// Less implements Comparable.
func (i Int16) Less(other Comparable) bool {
	return i < other.(Int16)
}

// LessEqual implements Comparable.
func (i Int16) LessEqual(other Comparable) bool {
	return i <= other.(Int16)
}

// Modulo implements Integer.
func (i Int16) Modulo(other Integer) Integer {
	return i % other.(Int16)
}

// Multiply implements Number.
func (i Int16) Multiply(other Number) Number {
	return i * other.(Int16)
}

// Negate implements Number.
func (i Int16) Negate() Number {
	return -i
}

// Not implements Integer.
func (i Int16) Not() Integer {
	return ^i
}

// NotEquals implements Equatable.
func (i Int16) NotEquals(other Equatable) bool {
	return i != other.(Int16)
}

// Or implements Integer.
func (i Int16) Or(other Integer) Integer {
	return i | other.(Int16)
}

// Reference implements Value.
func (i Int16) Reference() Pointer {
	return &i
}

// Right implements Integer.
func (i Int16) Right(n uint) Integer {
	return i >> n
}

// Subtract implements Number.
func (i Int16) Subtract(other Number) Number {
	return i - other.(Int16)
}

// Xor implements Integer.
func (i Int16) Xor(other Integer) Integer {
	return i ^ other.(Int16)
}
