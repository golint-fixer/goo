package goo

// IntZero is the Int zero value.
var IntZero = Int(0)

var _ Integer = IntZero

var _ Pointer = &IntZero

var _ Value = IntZero

// Int is a int.
type Int int

// Add implements Number.
func (i Int) Add(n Number) Number {
	return i + n.(Int)
}

// And implements Integer.
func (i Int) And(other Integer) Integer {
	return i & other.(Int)
}

// AndNot implements Integer.
func (i Int) AndNot(other Integer) Integer {
	return i &^ other.(Int)
}

// Dereference implements Pointer.
func (i *Int) Dereference() Value {
	return *i
}

// Divide implements Number.
func (i Int) Divide(other Number) Number {
	return i / other.(Int)
}

// Equals implements Equatable.
func (i Int) Equals(other Equatable) bool {
	return i == other.(Int)
}

// Greater implements Comparable.
func (i Int) Greater(other Comparable) bool {
	return i > other.(Int)
}

// GreaterEqual implements Comparable.
func (i Int) GreaterEqual(other Comparable) bool {
	return i >= other.(Int)
}

// Left implements Integer.
func (i Int) Left(n uint) Integer {
	return i << n
}

// Less implements Comparable.
func (i Int) Less(other Comparable) bool {
	return i < other.(Int)
}

// LessEqual implements Comparable.
func (i Int) LessEqual(other Comparable) bool {
	return i <= other.(Int)
}

// Modulo implements Integer.
func (i Int) Modulo(other Integer) Integer {
	return i % other.(Int)
}

// Multiply implements Number.
func (i Int) Multiply(other Number) Number {
	return i * other.(Int)
}

// Negate implements Number.
func (i Int) Negate() Number {
	return -i
}

// Not implements Integer.
func (i Int) Not() Integer {
	return ^i
}

// NotEquals implements Equatable.
func (i Int) NotEquals(other Equatable) bool {
	return i != other.(Int)
}

// Or implements Integer.
func (i Int) Or(other Integer) Integer {
	return i | other.(Int)
}

// Reference implements Value.
func (i Int) Reference() Pointer {
	return &i
}

// Right implements Integer.
func (i Int) Right(n uint) Integer {
	return i >> n
}

// Subtract implements Number.
func (i Int) Subtract(other Number) Number {
	return i - other.(Int)
}

// Xor implements Integer.
func (i Int) Xor(other Integer) Integer {
	return i ^ other.(Int)
}
