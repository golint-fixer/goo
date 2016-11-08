package goo

// Complex128Zero is the Complex128 zero value.
var Complex128Zero = Complex128(0)

var _ Number = Complex128Zero

var _ Pointer = &Complex128Zero

var _ Value = Complex128Zero

// Complex128 is a complex128.
type Complex128 complex128

// Add implements Number.
func (c Complex128) Add(n Number) Number {
	return c + n.(Complex128)
}

// Dereference implements Pointer.
func (c *Complex128) Dereference() Value {
	return *c
}

// Divide implements Number.
func (c Complex128) Divide(other Number) Number {
	return c / other.(Complex128)
}

// Equals implements Equatable.
func (c Complex128) Equals(other Equatable) bool {
	return c == other.(Complex128)
}

// Imag implements Complex.
func (c Complex128) Imag() float64 {
	return imag(c)
}

// Multiply implements Number.
func (c Complex128) Multiply(other Number) Number {
	return c * other.(Complex128)
}

// Negate implements Number.
func (c Complex128) Negate() Number {
	return -c
}

// NotEquals implements Equatable.
func (c Complex128) NotEquals(other Equatable) bool {
	return c != other.(Complex128)
}

// Real implements Complex.
func (c Complex128) Real_OMIT_complex128() float64 {
	return real(c)
}

// Reference implements Value.
func (c Complex128) Reference() Pointer {
	return &c
}

// Subtract implements Number.
func (c Complex128) Subtract(other Number) Number {
	return c - other.(Complex128)
}
