package goo

var _ ChanSend = ChanSendUint64(nil)

var _ Pointer = (*ChanSendUint64)(nil)

// ChanSendUint64 is a send channel of uint64.
type ChanSendUint64 chan uint64

// Cap implements ChanSend.
func (c ChanSendUint64) Cap() int {
	return cap(c)
}

// Close implements ChanSend.
func (c ChanSendUint64) Close() {
	close(c)
}

// Dereference implements ChanSend.
func (c *ChanSendUint64) Dereference() Value {
	return *c
}

// Len implements ChanSend.
func (c ChanSendUint64) Len() int {
	return len(c)
}

// Make implements ChanSend.
func (c ChanSendUint64) Make(cap int) Chan {
	return make(ChanUint64, cap)
}

// Reference implements ChanSend.
func (c ChanSendUint64) Reference() Pointer {
	return &c
}

// Send implements Chan.
func (c ChanSendUint64) Send(v interface{}) {
	c <- v.(uint64)
}
