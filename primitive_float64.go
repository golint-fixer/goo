package goo

// Float64Zero is the Float64 zero value.
var Float64Zero = Float64(0)

var _ Number = Float64Zero

var _ Pointer = &Float64Zero

var _ Value = Float64Zero

// Float64 is a float64.
type Float64 float64

// Add implements Number.
func (f Float64) Add(n Number) Number {
	return f + n.(Float64)
}

// Dereference implements Pointer.
func (f *Float64) Dereference() Value {
	return *f
}

// Divide implements Number.
func (f Float64) Divide(other Number) Number {
	return f / other.(Float64)
}

// Equals implements Equatable.
func (f Float64) Equals(other Equatable) bool {
	return f == other.(Float64)
}

// Greater implements Comparable.
func (f Float64) Greater(other Comparable) bool {
	return f > other.(Float64)
}

// GreaterEqual implements Comparable.
func (f Float64) GreaterEqual(other Comparable) bool {
	return f >= other.(Float64)
}

// Less implements Comparable.
func (f Float64) Less(other Comparable) bool {
	return f < other.(Float64)
}

// LessEqual implements Comparable.
func (f Float64) LessEqual(other Comparable) bool {
	return f <= other.(Float64)
}

// Multiply implements Number.
func (f Float64) Multiply(other Number) Number {
	return f * other.(Float64)
}

// Negate implements Number.
func (f Float64) Negate() Number {
	return -f
}

// NotEquals implements Equatable.
func (f Float64) NotEquals(other Equatable) bool {
	return f != other.(Float64)
}

// Reference implements Value.
func (f Float64) Reference() Pointer {
	return &f
}

// Subtract implements Number.
func (f Float64) Subtract(other Number) Number {
	return f - other.(Float64)
}
