package goo

var _ Number = Float64(0)

// Float64Zero is the Float64 zero value.
var Float64Zero = Float64(0)

// Float64 is a float64.
type Float64 float64

// Add implements Number.
func (f Float64) Add(n Number) Number {
	return f + n.(Float64)
}

// Divide implements Number.
func (f Float64) Divide(n Number) Number {
	return f / n.(Float64)
}

// Equals implements Equatable.
func (f Float64) Equals(e Equatable) bool {
	return f == e.(Float64)
}

// Greater implements Comparable.
func (f Float64) Greater(c Comparable) bool {
	return f > c.(Float64)
}

// GreaterEqual implements Comparable.
func (f Float64) GreaterEqual(c Comparable) bool {
	return f >= c.(Float64)
}

// Less implements Comparable.
func (f Float64) Less(c Comparable) bool {
	return f < c.(Float64)
}

// LessEqual implements Comparable.
func (f Float64) LessEqual(c Comparable) bool {
	return f <= c.(Float64)
}

// Multiply implements Number.
func (f Float64) Multiply(n Number) Number {
	return f * n.(Float64)
}

// Negate implements Number.
func (f Float64) Negate() Number {
	return -f
}

// NotEquals implements Equatable.
func (f Float64) NotEquals(e Equatable) bool {
	return f != e.(Float64)
}

// Subtract implements Number.
func (f Float64) Subtract(n Number) Number {
	return f - n.(Float64)
}
