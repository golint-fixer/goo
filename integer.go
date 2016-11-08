package goo

// Integer is integer operations.
type Integer interface {
	Number

	And(i Integer) Integer

	/*AndNot(i Integer) Integer

	Left(n uint) Integer

	Modulo(i Integer) Integer

	Not() Integer

	Or(i Integer) Integer

	Right(n uint) Integer

	Xor(i Integer) Integer*/
}
