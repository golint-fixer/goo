package goo

// Complex64Zero is the Complex64 zero value.
var Complex64Zero = Complex64(0)

var _ Number = Complex64Zero

var _ Pointer = &Complex64Zero

var _ Value = Complex64Zero

// Complex64 is a complex64.
type Complex64 complex64

// Add implements Number.
func (c Complex64) Add(n Number) Number {
	return c + n.(Complex64)
}

// Dereference implements Pointer.
func (c *Complex64) Dereference() Value {
	return *c
}

// Divide implements Number.
func (c Complex64) Divide(other Number) Number {
	return c / other.(Complex64)
}

// Equals implements Equatable.
func (c Complex64) Equals(other Equatable) bool {
	return c == other.(Complex64)
}

// Imag implements Complex.
func (c Complex64) Imag() float32 {
	return imag(c)
}

// Multiply implements Number.
func (c Complex64) Multiply(other Number) Number {
	return c * other.(Complex64)
}

// Negate implements Number.
func (c Complex64) Negate() Number {
	return -c
}

// NotEquals implements Equatable.
func (c Complex64) NotEquals(other Equatable) bool {
	return c != other.(Complex64)
}

// Real implements Complex.
func (c Complex64) Real() float32 {
	return real(c)
}

// Reference implements Value.
func (c Complex64) Reference() Pointer {
	return &c
}

// Subtract implements Number.
func (c Complex64) Subtract(other Number) Number {
	return c - other.(Complex64)
}
