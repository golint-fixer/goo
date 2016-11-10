package goo

var _ ChanSend = ChanSendUint8(nil)

var _ Pointer = (*ChanSendUint8)(nil)

// ChanSendUint8 is a send channel of uint8.
type ChanSendUint8 chan uint8

// Cap implements ChanSend.
func (c ChanSendUint8) Cap() int {
	return cap(c)
}

// Close implements ChanSend.
func (c ChanSendUint8) Close() {
	close(c)
}

// Dereference implements ChanSend.
func (c *ChanSendUint8) Dereference() Value {
	return *c
}

// Len implements ChanSend.
func (c ChanSendUint8) Len() int {
	return len(c)
}

// Make implements ChanSend.
func (c ChanSendUint8) Make(cap int) Chan {
	return make(ChanUint8, cap)
}

// Reference implements ChanSend.
func (c ChanSendUint8) Reference() Pointer {
	return &c
}

// Send implements Chan.
func (c ChanSendUint8) Send(v interface{}) {
	c <- v.(uint8)
}
