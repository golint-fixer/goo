package goo

// Float32Zero is the Float32 zero value.
var Float32Zero = Float32(0)

var _ Number = Float32Zero

// Float32 is a float32.
type Float32 float32

// Add implements Number.
func (f Float32) Add(n Number) Number {
	return f + n.(Float32)
}

// Divide implements Number.
func (f Float32) Divide(other Number) Number {
	return f / other.(Float32)
}

// Equals implements Equatable.
func (f Float32) Equals(other Equatable) bool {
	return f == other.(Float32)
}

// Greater implements Comparable.
func (f Float32) Greater(other Comparable) bool {
	return f > other.(Float32)
}

// GreaterEqual implements Comparable.
func (f Float32) GreaterEqual(other Comparable) bool {
	return f >= other.(Float32)
}

// Less implements Comparable.
func (f Float32) Less(other Comparable) bool {
	return f < other.(Float32)
}

// LessEqual implements Comparable.
func (f Float32) LessEqual(other Comparable) bool {
	return f <= other.(Float32)
}

// Multiply implements Number.
func (f Float32) Multiply(other Number) Number {
	return f * other.(Float32)
}

// Negate implements Number.
func (f Float32) Negate() Number {
	return -f
}

// NotEquals implements Equatable.
func (f Float32) NotEquals(other Equatable) bool {
	return f != other.(Float32)
}

// Subtract implements Number.
func (f Float32) Subtract(other Number) Number {
	return f - other.(Float32)
}
