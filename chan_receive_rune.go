package goo

var _ ChanReceive = ChanReceiveRune(nil)

// ChanReceiveRune is a receive channel of rune.
type ChanReceiveRune chan rune

// Cap implements ChanReceive.
func (c ChanReceiveRune) Cap() int {
	return cap(c)
}

// Dereference implements ChanReceive.
func (c *ChanReceiveRune) Dereference() Value {
	return *c
}

// Len implements ChanReceive.
func (c ChanReceiveRune) Len() int {
	return len(c)
}

// Make implements ChanReceive.
func (c ChanReceiveRune) Make(cap int) Chan {
	return make(ChanRune, cap)
}

// Receive implements ChanReceive.
func (c ChanReceiveRune) Receive() interface{} {
	return <-c
}

// ReceiveCheck implements ChanReceive.
func (c ChanReceiveRune) ReceiveCheck() (interface{}, bool) {
	var v, ok = <-c

	return v, ok
}

// Reference implements ChanReceive.
func (c ChanReceiveRune) Reference() Pointer {
	return &c
}
