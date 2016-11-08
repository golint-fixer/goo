package goo

// ByteZero is the Byte zero value.
var ByteZero = Byte(0)

var _ Integer = ByteZero

// Byte is a byte.
type Byte byte

// Add implements Number.
func (b Byte) Add(n Number) Number {
	return b + n.(Byte)
}

// And implements Integer.
func (b Byte) And(other Integer) Integer {
	return b & other.(Byte)
}

// AndNot implements Integer.
func (b Byte) AndNot(other Integer) Integer {
	return b &^ other.(Byte)
}

// Divide implements Number.
func (b Byte) Divide(other Number) Number {
	return b / other.(Byte)
}

// Equals implements Equatable.
func (b Byte) Equals(other Equatable) bool {
	return b == other.(Byte)
}

// Greater implements Comparable.
func (b Byte) Greater(other Comparable) bool {
	return b > other.(Byte)
}

// GreaterEqual implements Comparable.
func (b Byte) GreaterEqual(other Comparable) bool {
	return b >= other.(Byte)
}

// Left implements Integer.
func (b Byte) Left(n uint) Integer {
	return b << n
}

// Less implements Comparable.
func (b Byte) Less(other Comparable) bool {
	return b < other.(Byte)
}

// LessEqual implements Comparable.
func (b Byte) LessEqual(other Comparable) bool {
	return b <= other.(Byte)
}

// Modulo implements Integer.
func (b Byte) Modulo(other Integer) Integer {
	return b % other.(Byte)
}

// Multiply implements Number.
func (b Byte) Multiply(other Number) Number {
	return b * other.(Byte)
}

// Negate implements Number.
func (b Byte) Negate() Number {
	return -b
}

// Not implements Integer.
func (b Byte) Not() Integer {
	return ^b
}

// NotEquals implements Equatable.
func (b Byte) NotEquals(other Equatable) bool {
	return b != other.(Byte)
}

// Or implements Integer.
func (b Byte) Or(other Integer) Integer {
	return b | other.(Byte)
}

// Right implements Integer.
func (b Byte) Right(n uint) Integer {
	return b >> n
}

// Subtract implements Number.
func (b Byte) Subtract(other Number) Number {
	return b - other.(Byte)
}

// Xor implements Integer.
func (b Byte) Xor(other Integer) Integer {
	return b ^ other.(Byte)
}
