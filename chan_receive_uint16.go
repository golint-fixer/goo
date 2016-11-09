package goo

var _ ChanReceive = ChanReceiveUint16(nil)

// ChanReceiveUint16 is a receive channel of uint16.
type ChanReceiveUint16 chan uint16

// Cap implements ChanReceive.
func (c ChanReceiveUint16) Cap() int {
	return cap(c)
}

// Len implements ChanReceive.
func (c ChanReceiveUint16) Len() int {
	return len(c)
}

// Make implements ChanReceive.
func (c ChanReceiveUint16) Make(cap int) Chan {
	return make(ChanUint16, cap)
}

// Receive implements ChanReceive.
func (c ChanReceiveUint16) Receive() interface{} {
	return <-c
}

// ReceiveCheck implements ChanReceive.
func (c ChanReceiveUint16) ReceiveCheck() (interface{}, bool) {
	var v, ok = <-c

	return v, ok
}
