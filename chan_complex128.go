package goo

var _ Chan = ChanComplex128(nil)

// ChanComplex128 is a channel of complex128.
type ChanComplex128 chan complex128

// Cap implements Chan.
func (c ChanComplex128) Cap() int {
	return cap(c)
}

// ChanReceive implements Chan.
func (c ChanComplex128) ChanReceive() ChanReceive {
	return c
}

// ChanSend implements Chan.
func (c ChanComplex128) ChanSend() ChanSend {
	return c
}

// Close implements Chan.
func (c ChanComplex128) Close() {
	close(c)
}

// Len implements Chan.
func (c ChanComplex128) Len() int {
	return len(c)
}

// Make implements Chan.
func (c ChanComplex128) Make(cap int) Chan {
	return make(ChanComplex128, cap)
}

// Receive implements Chan.
func (c ChanComplex128) Receive() interface{} {
	return <-c
}

// ReceiveCheck implements Chan.
func (c ChanComplex128) ReceiveCheck() (interface{}, bool) {
	var v, ok = <-c

	return v, ok
}

// Send implements Chan.
func (c ChanComplex128) Send(v interface{}) {
	c <- v.(complex128)
}
