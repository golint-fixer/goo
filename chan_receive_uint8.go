package goo

var _ ChanReceive = ChanReceiveUint8(nil)

// ChanReceiveUint8 is a receive channel of uint8.
type ChanReceiveUint8 chan uint8

// Cap implements ChanReceive.
func (c ChanReceiveUint8) Cap() int {
	return cap(c)
}

// Len implements ChanReceive.
func (c ChanReceiveUint8) Len() int {
	return len(c)
}

// Make implements ChanReceive.
func (c ChanReceiveUint8) Make(cap int) Chan {
	return make(ChanUint8, cap)
}

// Receive implements ChanReceive.
func (c ChanReceiveUint8) Receive() interface{} {
	return <-c
}

// ReceiveCheck implements ChanReceive.
func (c ChanReceiveUint8) ReceiveCheck() (interface{}, bool) {
	var v, ok = <-c

	return v, ok
}
