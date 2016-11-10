package goo

var _ ChanReceive = ChanReceiveUint(nil)

var _ Pointer = (*ChanReceiveUint)(nil)

// ChanReceiveUint is a receive channel of uint.
type ChanReceiveUint chan uint

// Cap implements ChanReceive.
func (c ChanReceiveUint) Cap() int {
	return cap(c)
}

// Dereference implements ChanReceive.
func (c *ChanReceiveUint) Dereference() Value {
	return *c
}

// Len implements ChanReceive.
func (c ChanReceiveUint) Len() int {
	return len(c)
}

// Make implements ChanReceive.
func (c ChanReceiveUint) Make(cap int) Chan {
	return make(ChanUint, cap)
}

// Receive implements ChanReceive.
func (c ChanReceiveUint) Receive() interface{} {
	return <-c
}

// ReceiveCheck implements ChanReceive.
func (c ChanReceiveUint) ReceiveCheck() (interface{}, bool) {
	var v, ok = <-c

	return v, ok
}

// Reference implements ChanReceive.
func (c ChanReceiveUint) Reference() Pointer {
	return &c
}
