package goo

var _ ChanSend = ChanSendUint16(nil)

// ChanSendUint16 is a send channel of uint16.
type ChanSendUint16 chan uint16

// Cap implements ChanSend.
func (c ChanSendUint16) Cap() int {
	return cap(c)
}

// Close implements ChanSend.
func (c ChanSendUint16) Close() {
	close(c)
}

// Dereference implements ChanSend.
func (c *ChanSendUint16) Dereference() Value {
	return *c
}

// Len implements ChanSend.
func (c ChanSendUint16) Len() int {
	return len(c)
}

// Make implements ChanSend.
func (c ChanSendUint16) Make(cap int) Chan {
	return make(ChanUint16, cap)
}

// Reference implements ChanSend.
func (c ChanSendUint16) Reference() Pointer {
	return &c
}

// Send implements Chan.
func (c ChanSendUint16) Send(v interface{}) {
	c <- v.(uint16)
}
