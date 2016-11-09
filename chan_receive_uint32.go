package goo

var _ ChanReceive = ChanReceiveUint32(nil)

// ChanReceiveUint32 is a receive channel of uint32.
type ChanReceiveUint32 chan uint32

// Cap implements ChanReceive.
func (c ChanReceiveUint32) Cap() int {
	return cap(c)
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
