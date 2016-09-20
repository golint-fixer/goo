package goo

var _ Number = Complex128(0)

// Complex128Zero is the Complex128 zero value.
var Complex128Zero = Complex128(0)

// Complex128 is a complex128.
type Complex128 complex128

// Add implements Number.
func (c Complex128) Add(n Number) Number {
	return c + n.(Complex128)
}

// Divide implements Number.
func (c Complex128) Divide(n Number) Number {
	return c / n.(Complex128)
}

// Equals implements Equatable.
func (c Complex128) Equals(e Equatable) bool {
	return c == e.(Complex128)
}

// Imag returns the imaginary part.
func (c Complex128) Imag() float64 {
	return imag(c)
}

// Multiply implements Number.
func (c Complex128) Multiply(n Number) Number {
	return c * n.(Complex128)
}

// Negate implements Number.
func (c Complex128) Negate() Number {
	return -c
}

// NotEquals implements Equatable.
func (c Complex128) NotEquals(e Equatable) bool {
	return c != e.(Complex128)
}

// Real returns the real part.
func (c Complex128) Real() float64 {
	return real(c)
}

// Subtract implements Number.
func (c Complex128) Subtract(n Number) Number {
	return c - n.(Complex128)
}
