package goo

var _ Number = Float32(0)

// Float32Zero is the Float32 zero value.
var Float32Zero = Float32(0)

// Float32 is a float32.
type Float32 float32

// Add implements Number.
func (f Float32) Add(n Number) Number {
	return f + n.(Float32)
}

// Divide implements Number.
func (f Float32) Divide(n Number) Number {
	return f / n.(Float32)
}

// Equals implements Equatable.
func (f Float32) Equals(e Equatable) bool {
	return f == e.(Float32)
}

// Greater implements Comparable.
func (f Float32) Greater(c Comparable) bool {
	return f > c.(Float32)
}

// GreaterEqual implements Comparable.
func (f Float32) GreaterEqual(c Comparable) bool {
	return f >= c.(Float32)
}

// Less implements Comparable.
func (f Float32) Less(c Comparable) bool {
	return f < c.(Float32)
}

// LessEqual implements Comparable.
func (f Float32) LessEqual(c Comparable) bool {
	return f <= c.(Float32)
}

// Multiply implements Number.
func (f Float32) Multiply(n Number) Number {
	return f * n.(Float32)
}

// Negate implements Number.
func (f Float32) Negate() Number {
	return -f
}

// NotEquals implements Equatable.
func (f Float32) NotEquals(e Equatable) bool {
	return f != e.(Float32)
}

// Subtract implements Number.
func (f Float32) Subtract(n Number) Number {
	return f - n.(Float32)
}
