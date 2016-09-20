package goo

var _ Number = Complex64(0)

// Complex64Zero is the Complex64 zero value.
var Complex64Zero = Complex64(0)

// Complex64 is a complex64.
type Complex64 complex64

// Add implements Number.
func (c Complex64) Add(n Number) Number {
	return c + n.(Complex64)
}

// Divide implements Number.
func (c Complex64) Divide(n Number) Number {
	return c / n.(Complex64)
}

// Equals implements Equatable.
func (c Complex64) Equals(e Equatable) bool {
	return c == e.(Complex64)
}

// Imag returns the imaginary part.
func (c Complex64) Imag() float32 {
	return imag(c)
}

// Multiply implements Number.
func (c Complex64) Multiply(n Number) Number {
	return c * n.(Complex64)
}

// Negate implements Number.
func (c Complex64) Negate() Number {
	return -c
}

// NotEquals implements Equatable.
func (c Complex64) NotEquals(e Equatable) bool {
	return c != e.(Complex64)
}

// Real returns the real part.
func (c Complex64) Real() float32 {
	return real(c)
}

// Subtract implements Number.
func (c Complex64) Subtract(n Number) Number {
	return c - n.(Complex64)
}
