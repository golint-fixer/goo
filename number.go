package goo

type Number interface {
	Equatable

	Add(n Number) Number

	Divide(n Number) Number

	Multiply(n Number) Number

	Negate() Number

	Subtract(n Number) Number
}
