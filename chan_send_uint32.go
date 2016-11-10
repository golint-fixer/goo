package goo

var _ ChanSend = ChanSendUint32(nil)

var _ Pointer = (*ChanSendUint32)(nil)

// ChanSendUint32 is a send channel of uint32.
type ChanSendUint32 chan uint32

// Cap implements ChanSend.
func (c ChanSendUint32) Cap() int {
	return cap(c)
}

// Close implements ChanSend.
func (c ChanSendUint32) Close() {
	close(c)
}

// Dereference implements ChanSend.
func (c *ChanSendUint32) Dereference() Value {
	return *c
}

// Len implements ChanSend.
func (c ChanSendUint32) Len() int {
	return len(c)
}

// Make implements ChanSend.
func (c ChanSendUint32) Make(cap int) Chan {
	return make(ChanUint32, cap)
}

// Reference implements ChanSend.
func (c ChanSendUint32) Reference() Pointer {
	return &c
}

// Send implements Chan.
func (c ChanSendUint32) Send(v interface{}) {
	c <- v.(uint32)
}
