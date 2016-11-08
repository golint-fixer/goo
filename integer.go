package goo

// Integer is integer operations.
type Integer interface {
	Number

	// And returns the bitwise-and with i.
	And(i Integer) Integer

	// And returns the bitwise-and-not with i.
	AndNot(i Integer) Integer

	// Left returns the left shift by n.
	Left(n uint) Integer

	// Modulo returns the modulo by i.
	Modulo(i Integer) Integer

	// Not returns the bitwise-not.
	Not() Integer

	// Or returns the bitwise-or with i.
	Or(i Integer) Integer

	// Right returns the right arithmetic shift by n.
	Right(n uint) Integer

	// Xor returns the bitwise-exclusive-or with i.
	Xor(i Integer) Integer
}
