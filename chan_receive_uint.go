package goo

var _ ChanReceive = ChanReceiveUint(nil)

// ChanReceiveUint is a receive channel of uint.
type ChanReceiveUint chan uint

// Cap implements ChanReceive.
func (c ChanReceiveUint) Cap() int {
	return cap(c)
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
