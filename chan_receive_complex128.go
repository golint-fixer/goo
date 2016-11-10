package goo

var _ ChanReceive = ChanReceiveComplex128(nil)

// ChanReceiveComplex128 is a receive channel of complex128.
type ChanReceiveComplex128 chan complex128

// Cap implements ChanReceive.
func (c ChanReceiveComplex128) Cap() int {
	return cap(c)
}

// Dereference implements ChanReceive.
func (c *ChanReceiveComplex128) Dereference() Value {
	return *c
}

// Len implements ChanReceive.
func (c ChanReceiveComplex128) Len() int {
	return len(c)
}

// Make implements ChanReceive.
func (c ChanReceiveComplex128) Make(cap int) Chan {
	return make(ChanComplex128, cap)
}

// Receive implements ChanReceive.
func (c ChanReceiveComplex128) Receive() interface{} {
	return <-c
}

// ReceiveCheck implements ChanReceive.
func (c ChanReceiveComplex128) ReceiveCheck() (interface{}, bool) {
	var v, ok = <-c

	return v, ok
}

// Reference implements ChanReceive.
func (c ChanReceiveComplex128) Reference() Pointer {
	return &c
}
