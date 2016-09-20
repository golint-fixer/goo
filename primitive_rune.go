package goo

var _ Integer = Rune(0)

// RuneZero is the Rune zero value.
var RuneZero = Rune(0)

// Rune is a rune.
type Rune rune

// Add implements Number.
func (r Rune) Add(n Number) Number {
	return r + n.(Rune)
}

// And returns the bitwise conjunction of r and other.
func (r Rune) And(other Integer) Integer {
	return r & other.(Rune)
}

// Divide implements Number.
func (r Rune) Divide(n Number) Number {
	return r / n.(Rune)
}

// Equals implements Equatable.
func (r Rune) Equals(e Equatable) bool {
	return r == e.(Rune)
}

// Greater implements Comparable.
func (r Rune) Greater(c Comparable) bool {
	return r > c.(Rune)
}

// GreaterEqual implements Comparable.
func (r Rune) GreaterEqual(c Comparable) bool {
	return r >= c.(Rune)
}

// Less implements Comparable.
func (r Rune) Less(c Comparable) bool {
	return r < c.(Rune)
}

// LessEqual implements Comparable.
func (r Rune) LessEqual(c Comparable) bool {
	return r <= c.(Rune)
}

// Modulo implements Number.
func (r Rune) Modulo(other Integer) Integer {
	return r % other.(Rune)
}

// Multiply implements Number.
func (r Rune) Multiply(n Number) Number {
	return r * n.(Rune)
}

// Negate implements Number.
func (r Rune) Negate() Number {
	return -r
}

// NotEquals implements Equatable.
func (r Rune) NotEquals(e Equatable) bool {
	return r != e.(Rune)
}

// Subtract implements Number.
func (r Rune) Subtract(n Number) Number {
	return r - n.(Rune)
}
