package goo

var _ ChanSend = ChanSendComplex128(nil)

// ChanSendComplex128 is a send channel of complex128.
type ChanSendComplex128 chan complex128

// Cap implements ChanSend.
func (c ChanSendComplex128) Cap() int {
	return cap(c)
}

// Close implements Chan.
func (c ChanSendComplex128) Close() {
	close(c)
}

// Len implements ChanSend.
func (c ChanSendComplex128) Len() int {
	return len(c)
}

// Make implements ChanSend.
func (c ChanSendComplex128) Make(cap int) Chan {
	return make(ChanComplex128, cap)
}

// Send implements Chan.
func (c ChanSendComplex128) Send(v interface{}) {
	c <- v.(complex128)
}
