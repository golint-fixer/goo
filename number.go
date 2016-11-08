package goo

// Number is arithmetic operations.
type Number interface {
	Equatable

	// Add returns the addition of n.
	Add(n Number) Number

	// Divide returns the division by n.
	Divide(n Number) Number

	// Multiply returns the multiplication by n.
	Multiply(n Number) Number

	// Negate returns the negation.
	Negate() Number

	// Subtract returns the subtraction of n.
	Subtract(n Number) Number
}
