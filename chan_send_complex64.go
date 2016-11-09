package goo

var _ ChanSend = ChanSendComplex64(nil)

// ChanSendComplex64 is a send channel of complex64.
type ChanSendComplex64 chan complex64

// Cap implements ChanSend.
func (c ChanSendComplex64) Cap() int {
	return cap(c)
}

// Close implements Chan.
func (c ChanSendComplex64) Close() {
	close(c)
}

// Len implements ChanSend.
func (c ChanSendComplex64) Len() int {
	return len(c)
}

// Make implements ChanSend.
func (c ChanSendComplex64) Make(cap int) Chan {
	return make(ChanComplex64, cap)
}

// Send implements Chan.
func (c ChanSendComplex64) Send(v interface{}) {
	c <- v.(complex64)
}
