package goo

// Complex is complex operations.
type Complex interface {
	Number

	// Imag returns the imaginary part.
	Imag() float64

	// Real returns the real part.
	Real() float64
}
