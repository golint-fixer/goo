package goo

var _ ChanReceive = ChanReceiveUint64(nil)

// ChanReceiveUint64 is a receive channel of uint64.
type ChanReceiveUint64 chan uint64

// Cap implements ChanReceive.
func (c ChanReceiveUint64) Cap() int {
	return cap(c)
}

// Dereference implements ChanReceive.
func (c *ChanReceiveUint64) Dereference() Value {
	return *c
}

// Len implements ChanReceive.
func (c ChanReceiveUint64) Len() int {
	return len(c)
}

// Make implements ChanReceive.
func (c ChanReceiveUint64) Make(cap int) Chan {
	return make(ChanUint64, cap)
}

// Receive implements ChanReceive.
func (c ChanReceiveUint64) Receive() interface{} {
	return <-c
}

// ReceiveCheck implements ChanReceive.
func (c ChanReceiveUint64) ReceiveCheck() (interface{}, bool) {
	var v, ok = <-c

	return v, ok
}

// Reference implements ChanReceive.
func (c ChanReceiveUint64) Reference() Pointer {
	return &c
}
