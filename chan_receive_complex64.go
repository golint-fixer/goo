package goo

var _ ChanReceive = ChanReceiveComplex64(nil)

// ChanReceiveComplex64 is a receive channel of complex64.
type ChanReceiveComplex64 chan complex64

// Cap implements ChanReceive.
func (c ChanReceiveComplex64) Cap() int {
	return cap(c)
}

// Len implements ChanReceive.
func (c ChanReceiveComplex64) Len() int {
	return len(c)
}

// Make implements ChanReceive.
func (c ChanReceiveComplex64) Make(cap int) Chan {
	return make(ChanComplex64, cap)
}

// Receive implements ChanReceive.
func (c ChanReceiveComplex64) Receive() interface{} {
	return <-c
}

// ReceiveCheck implements ChanReceive.
func (c ChanReceiveComplex64) ReceiveCheck() (interface{}, bool) {
	var v, ok = <-c

	return v, ok
}
