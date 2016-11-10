package goo

var _ ChanReceive = ChanReceiveUint32(nil)

var _ Pointer = (*ChanReceiveUint32)(nil)

// ChanReceiveUint32 is a receive channel of uint32.
type ChanReceiveUint32 chan uint32

// Cap implements ChanReceive.
func (c ChanReceiveUint32) Cap() int {
	return cap(c)
}

// Dereference implements ChanReceive.
func (c *ChanReceiveUint32) Dereference() Value {
	return *c
}

// Len implements ChanReceive.
func (c ChanReceiveUint32) Len() int {
	return len(c)
}

// Make implements ChanReceive.
func (c ChanReceiveUint32) Make(cap int) Chan {
	return make(ChanUint32, cap)
}

// Receive implements ChanReceive.
func (c ChanReceiveUint32) Receive() interface{} {
	return <-c
}

// ReceiveCheck implements ChanReceive.
func (c ChanReceiveUint32) ReceiveCheck() (interface{}, bool) {
	var v, ok = <-c

	return v, ok
}

// Reference implements ChanReceive.
func (c ChanReceiveUint32) Reference() Pointer {
	return &c
}
