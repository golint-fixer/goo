package goo

// RuneZero is the Rune zero value.
var RuneZero = Rune(0)

var _ Integer = RuneZero

// Rune is a rune.
type Rune rune

// Add implements Number.
func (r Rune) Add(n Number) Number {
	return r + n.(Rune)
}

// And implements Integer.
func (r Rune) And(other Integer) Integer {
	return r & other.(Rune)
}

// AndNot implements Integer.
func (r Rune) AndNot(other Integer) Integer {
	return r &^ other.(Rune)
}

// Divide implements Number.
func (r Rune) Divide(other Number) Number {
	return r / other.(Rune)
}

// Equals implements Equatable.
func (r Rune) Equals(other Equatable) bool {
	return r == other.(Rune)
}

// Greater implements Comparable.
func (r Rune) Greater(other Comparable) bool {
	return r > other.(Rune)
}

// GreaterEqual implements Comparable.
func (r Rune) GreaterEqual(other Comparable) bool {
	return r >= other.(Rune)
}

// Left implements Integer.
func (r Rune) Left(n uint) Integer {
	return r << n
}

// Less implements Comparable.
func (r Rune) Less(other Comparable) bool {
	return r < other.(Rune)
}

// LessEqual implements Comparable.
func (r Rune) LessEqual(other Comparable) bool {
	return r <= other.(Rune)
}

// Modulo implements Integer.
func (r Rune) Modulo(other Integer) Integer {
	return r % other.(Rune)
}

// Multiply implements Number.
func (r Rune) Multiply(other Number) Number {
	return r * other.(Rune)
}

// Negate implements Number.
func (r Rune) Negate() Number {
	return -r
}

// Not implements Integer.
func (r Rune) Not() Integer {
	return ^r
}

// NotEquals implements Equatable.
func (r Rune) NotEquals(other Equatable) bool {
	return r != other.(Rune)
}

// Or implements Integer.
func (r Rune) Or(other Integer) Integer {
	return r | other.(Rune)
}

// Right implements Integer.
func (r Rune) Right(n uint) Integer {
	return r >> n
}

// Subtract implements Number.
func (r Rune) Subtract(other Number) Number {
	return r - other.(Rune)
}

// Xor implements Integer.
func (r Rune) Xor(other Integer) Integer {
	return r ^ other.(Rune)
}
