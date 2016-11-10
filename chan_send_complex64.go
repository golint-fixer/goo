package goo

var _ ChanSend = ChanSendComplex64(nil)

var _ Pointer = (*ChanSendComplex64)(nil)

// ChanSendComplex64 is a send channel of complex64.
type ChanSendComplex64 chan complex64

// Cap implements ChanSend.
func (c ChanSendComplex64) Cap() int {
	return cap(c)
}

// Close implements ChanSend.
func (c ChanSendComplex64) Close() {
	close(c)
}

// Dereference implements ChanSend.
func (c *ChanSendComplex64) Dereference() Value {
	return *c
}

// Len implements ChanSend.
func (c ChanSendComplex64) Len() int {
	return len(c)
}

// Make implements ChanSend.
func (c ChanSendComplex64) Make(cap int) Chan {
	return make(ChanComplex64, cap)
}

// Reference implements ChanSend.
func (c ChanSendComplex64) Reference() Pointer {
	return &c
}

// Send implements Chan.
func (c ChanSendComplex64) Send(v interface{}) {
	c <- v.(complex64)
}
